package main

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"errors"
	"flag"
	"fmt"
	"github.com/bsthun/gut"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	config.BootConfiguration()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetInt("DB_PORT"),
	)
	lg := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: lg,
	})
	if err != nil {
		gut.Fatal("Failed to connect to database", err)
	}

	parentDocumentId := flag.String("parentDocumentId", "", "Outline document ID")
	documentId := flag.String("documentId", "", "Outline document ID")
	flag.Parse()

	if *parentDocumentId == "" {
		gut.Fatal("missing required flag: parentDocumentId", nil)
	}

	if *documentId != "" {
		processDocument(db, documentId)
	}

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(*config.Env.OutlineToken).
		SetBody(map[string]any{
			"id": *parentDocumentId,
		}).
		SetResult(map[string]any{}).
		Post("https://outline.cscms.me/api/documents.info")
	if err != nil {
		gut.Fatal("failed to call outline api", err)
	}

	*parentDocumentId = (*resp.Result().(*map[string]any))["data"].(map[string]any)["id"].(string)

	resp, err = client.R().
		SetAuthToken(*config.Env.OutlineToken).
		SetBody(map[string]any{
			"parentDocumentId": *parentDocumentId,
		}).
		SetResult(map[string]any{}).
		Post("https://outline.cscms.me/api/documents.list")
	if err != nil {
		gut.Fatal("failed to call outline api", err)
	}

	documents := (*resp.Result().(*map[string]any))["data"].([]any)
	for _, document := range documents {
		documentId := document.(map[string]any)["id"].(string)
		processDocument(db, &documentId)
	}
}

type CourseMetadata struct {
	Name        string
	ImageUrl    string
	Description string
	FieldName   string
}

func processDocument(db *gorm.DB, documentId *string) {
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(*config.Env.OutlineToken).
		SetBody(map[string]any{
			"id": documentId,
		}).
		SetResult(map[string]any{}).
		Post("https://outline.cscms.me/api/documents.export")
	if err != nil {
		gut.Fatal("failed to call outline api", err)
	}

	markdown := (*resp.Result().(*map[string]any))["data"].(string)
	lines := strings.Split(markdown, "\n")
	if len(lines) < 1 {
		gut.Fatal("malformed markdown: empty content", nil)
	}

	metadata := extractCourseMetadata(lines)
	course := createOrFindCourse(db, &metadata)
	processContent(db, course, lines)
}

func extractCourseMetadata(lines []string) CourseMetadata {
	var metadata CourseMetadata
	var currentSection string

	metadata.Name = strings.TrimPrefix(strings.TrimSpace(lines[0]), "# ")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "## ") {
			currentSection = strings.TrimPrefix(line, "## ")
			continue
		}

		if line == "" {
			continue
		}

		switch currentSection {
		case "Image":
			metadata.ImageUrl = strings.Trim(line, "<>")
		case "Description":
			metadata.Description = line
		case "Field":
			metadata.FieldName = line
		}
	}

	return metadata
}

func createOrFindCourse(db *gorm.DB, metadata *CourseMetadata) *models.Course {
	var field models.FieldType
	if err := db.Where("name = ?", metadata.FieldName).First(&field).Error; err != nil {
		gut.Fatal("field not found: "+metadata.FieldName, err)
	}

	var course *models.Course
	tx := db.Where("name = ?", metadata.Name).First(&course)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		course = &models.Course{
			Name:    &metadata.Name,
			FieldId: field.Id,
		}
		if err := db.Create(&course).Error; err != nil {
			gut.Fatal("failed to create course", err)
		}
	}

	return course
}

func processContent(db *gorm.DB, course *models.Course, lines []string) {
	lines = lines[1:]
	var currentSection string
	var contentOrder int64 = 1
	var contentBuffer strings.Builder
	var moduleTitle string
	inMetadataSection := false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "## ") {
			section := strings.TrimPrefix(line, "## ")
			inMetadataSection = section == "Image" || section == "Description" || section == "Field"

			if !inMetadataSection && contentBuffer.Len() > 0 {
				saveContent(db, course, &contentOrder, currentSection, contentBuffer.String(), moduleTitle)
				contentBuffer.Reset()
			}
			currentSection = section
			continue
		}

		if line == "" || inMetadataSection {
			continue
		}

		if currentSection == "Module" {
			moduleTitle = line
			saveContent(db, course, &contentOrder, currentSection, "", line)
		} else {
			contentBuffer.WriteString(line + "\n")
		}
	}

	if !inMetadataSection && contentBuffer.Len() > 0 {
		saveContent(db, course, &contentOrder, currentSection, contentBuffer.String(), moduleTitle)
	}
}

func saveContent(db *gorm.DB, course *models.Course, order *int64, contentType, content, moduleTitle string) {
	var moduleId *uint64
	if contentType == "Module" {
		var module models.Module
		if err := db.Where("title = ?", moduleTitle).First(&module).Error; err != nil {
			gut.Fatal("module not found: "+moduleTitle, err)
		}
		moduleId = module.Id
	}

	var contentTypeStr string
	if contentType == "Module" {
		contentTypeStr = "module"
	} else {
		contentTypeStr = "text"
	}

	var textPtr *string
	if contentTypeStr != "module" {
		textPtr = gut.Ptr(content)
	}

	courseContent := &models.CourseContent{
		CourseId: course.Id,
		Order:    order,
		Type:     &contentTypeStr,
		Text:     textPtr,
		ModuleId: moduleId,
	}

	if err := db.Create(&courseContent).Error; err != nil {
		gut.Fatal("failed to create course content", err)
	}

	*order++
}
