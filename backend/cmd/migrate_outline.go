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
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// initialize config
	config.BootConfiguration()

	// connect to database
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

	// parse flags
	documentId := flag.String("documentId", "", "Outline document ID")
	flag.Parse()

	// validate flags
	if *documentId == "" {
		gut.Fatal("missing required flag: documentId", nil)
	}

	// call outline api
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(*config.Env.OutlineToken).
		SetHeader("Accept", "*/*").
		SetBody(map[string]string{
			"id": *documentId,
		}).
		SetResult(map[string]any{}).
		Post("https://outline.cscms.me/api/documents.export")
	if err != nil {
		gut.Fatal("failed to call outline api", err)
	}

	// get markdown content
	markdown := (*resp.Result().(*map[string]interface{}))["data"].(string)

	// split content into sections
	sections := strings.Split(markdown, "\n# ")
	if len(sections) < 2 {
		gut.Fatal("malformed markdown: missing module", nil)
	}

	// process module title
	moduleTitle := strings.Split(sections[0], "\n")[0]
	moduleTitle = strings.TrimPrefix(moduleTitle, "# ")

	// construct module
	var module *models.Module

	// find or create module
	tx := db.Where("title = ?", moduleTitle).First(&module)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		module = &models.Module{
			Title:       &moduleTitle,
			ImageUrl:    new(string),
			Description: new(string),
		}
		if err := db.Create(&module).Error; err != nil {
			gut.Fatal("failed to create module", err)
		}
	}

	// process each step section
	for _, section := range sections[1:] {
		lines := strings.Split(section, "\n")
		if len(lines) < 1 {
			continue
		}

		stepTitle := strings.TrimSpace(lines[0])
		stepTitleParts := strings.Split(stepTitle, ": ")
		stepTitle = stepTitleParts[len(stepTitleParts)-1]
		var step *models.Step

		// find or create step
		result := db.Where("module_id = ? AND title = ?", module.Id, stepTitle).First(&step)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			step = &models.Step{
				Id:          nil,
				ModuleId:    module.Id,
				Module:      nil,
				Title:       &stepTitle,
				Description: nil,
				Content:     nil,
				Outcome:     nil,
				Check:       nil,
				Error:       nil,
				CreatedAt:   nil,
				UpdatedAt:   nil,
			}
			if tx := db.Create(step); tx.Error != nil {
				gut.Fatal("failed to create step", tx.Error)
			}
		}

		// 6. Update step content
		var description, content, outcome, check, errorable string

		currentSection := ""
		var evalBuffer []string

		for _, line := range lines[1:] {
			if strings.HasPrefix(line, "## ") {
				currentSection = strings.TrimSpace(strings.TrimPrefix(line, "## "))
				continue
			}

			switch currentSection {
			case "Description":
				description += line + "\n"
			case "Content":
				content += line + "\n"
			case "Outcome":
				outcome += line + "\n"
			case "Check":
				check += line + "\n"
			case "Error":
				errorable += line + "\n"
			case "Evaluation":
				if strings.HasPrefix(strings.TrimSpace(line), "* ") {
					evalBuffer = append(evalBuffer, strings.TrimPrefix(strings.TrimSpace(line), "* "))
				}
			}
		}

		// 7. Replace attachments paths
		description = replaceAttachmentPaths(content)
		content = replaceAttachmentPaths(content)
		outcome = replaceAttachmentPaths(outcome)
		check = replaceAttachmentPaths(check)
		errorable = replaceAttachmentPaths(errorable)

		// Verify required sections
		if description == "" || content == "" || outcome == "" || check == "" || errorable == "" {
			gut.Fatal("malformed markdown: missing required sections", nil)
		}

		// Update step
		step.Description = &description
		step.Content = &content
		step.Outcome = &outcome
		step.Check = &check
		step.Error = &errorable

		if err := db.Save(&step).Error; err != nil {
			gut.Fatal("failed to update step", err)
		}

		for i := 0; i < len(evalBuffer); i += 4 {
			if i+1 >= len(evalBuffer) {
				gut.Fatal("malformed markdown: missing evaluation type", nil)
			}

			evalType := evalBuffer[i+1]
			if evalType != "check" && evalType != "text" && evalType != "image" {
				gut.Fatal("malformed markdown: invalid evaluation type", nil)
			}

			instruction := evalBuffer[i+2]
			gem, _ := strconv.ParseInt(evalBuffer[i+3], 10, 64)

			order := i/4 + 1

			// Check if an entry with the same Order exists
			var existingEvaluation models.StepEvaluate
			if err := db.Where("step_id = ? AND \"order\" = ?", step.Id, order).First(&existingEvaluation).Error; err == nil {
				// Update the existing entry
				existingEvaluation.Question = &evalBuffer[i]
				existingEvaluation.Type = &evalType
				existingEvaluation.Instruction = &instruction
				existingEvaluation.Gem = gut.Ptr(int(gem))
				if err := db.Save(&existingEvaluation).Error; err != nil {
					gut.Fatal("failed to update evaluation", err)
				}
			} else {
				// Create a new entry
				evaluation := &models.StepEvaluate{
					Id:          nil,
					StepId:      step.Id,
					Step:        nil,
					Gem:         gut.Ptr(int(gem)),
					Order:       gut.Ptr(order),
					Question:    &evalBuffer[i],
					Type:        &evalType,
					Instruction: &instruction,
					CreatedAt:   nil,
					UpdatedAt:   nil,
				}

				if err := db.Create(&evaluation).Error; err != nil {
					gut.Fatal("failed to create evaluation", err)
				}
			}
		}
	}
}

func replaceAttachmentPaths(content string) string {
	re := regexp.MustCompile(`attachments/([^)]+\.png)`)
	content = strings.TrimSpace(content)
	return re.ReplaceAllString(content, "/prefix/$1")
}
