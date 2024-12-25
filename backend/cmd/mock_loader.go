package main

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"backend/internals/migration"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/bsthun/gut"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MockData struct {
	Users              []map[string]interface{} `json:"users"`
	FieldTypes         []map[string]interface{} `json:"field_types"`
	Courses            []map[string]interface{} `json:"courses"`
	Modules            []map[string]interface{} `json:"modules"`
	Steps              []map[string]interface{} `json:"steps"`
	StepAuthors        []map[string]interface{} `json:"step_authors"`
	StepEvaluates      []map[string]interface{} `json:"step_evaluates"`
	UserEvaluates      []map[string]interface{} `json:"user_evaluates"`
	StepComments       []map[string]interface{} `json:"step_comments"`
	StepCommentUpvotes []map[string]interface{} `json:"step_comment_upvotes"`
	UserPasses         []map[string]interface{} `json:"user_passes"`
	Enrolls            []map[string]interface{} `json:"enrolls"`
	CourseContents     []map[string]interface{} `json:"course_contents"`
	Articles           []map[string]interface{} `json:"articles"`
	UserActivity       []map[string]interface{} `json:"user_activity"`
}

func main() {

	config.BootConfiguration()

	// Read and parse mock data
	basePath, _ := os.Getwd()
	mockDataPath := filepath.Join(basePath, "mockData.json")

	mockDataBytes, err := os.ReadFile(mockDataPath)
	if err != nil {
		gut.Fatal("Failed to read mock data file", err)
	}

	var mockData MockData
	if err := json.Unmarshal(mockDataBytes, &mockData); err != nil {
		gut.Fatal("Failed to parse mock data", err)
	}

	lg := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetInt("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: lg,
	})
	if err != nil {
		gut.Fatal("Failed to connect to database", err)
	}

	// Clean existing tables
	var tables []string
	if tx := db.Raw(`SELECT tablename FROM pg_tables WHERE schemaname = 'public'`).Scan(&tables); tx.Error != nil {
		gut.Fatal("Failed to get tables", tx.Error)
	}
	for _, t := range tables {
		if tx := db.Exec("DROP TABLE IF EXISTS " + t + " CASCADE"); tx.Error != nil {
			gut.Fatal("Failed to drop table "+t, tx.Error)
		}
	}

	if viper.GetBool("DB_AUTOMIGRATE") {

		if err := db.AutoMigrate(
			&models.Article{},
			&models.Course{},
			&models.CourseContent{},
			&models.Enroll{},
			&models.FieldType{},
			&models.Module{},
			&models.Step{},
			&models.StepAuthor{},
			&models.StepComment{},
			&models.StepCommentUpvote{},
			&models.StepEvaluate{},
			&models.User{},
			&models.UserEvaluate{},
			&models.UserPass{},
			&models.UserActivity{},
		); err != nil {
			gut.Fatal("Failed to migrate schema", err)
		}
	}

	gut.Debug("Starting data migration...")

	if err := migration.MigrateUsers(db, mockData.Users); err != nil {
		gut.Error("failed to migrate users: %w", err)
	}

	if err := migration.MigrateFieldTypes(db, mockData.FieldTypes); err != nil {
		gut.Error("failed to migrate fields: %w", err)
	}

	if err := migration.MigrateModules(db, mockData.Modules); err != nil {
		gut.Error("failed to migrate modules: %w", err)
	}

	if err := migration.MigrateArticles(db, mockData.Articles); err != nil {
		gut.Error("failed to migrate articles: %w", err)
	}

	if err := migration.MigrateCourses(db, mockData.Courses); err != nil {
		gut.Error("failed to migrate courses: %w", err)
	}

	if err := migration.MigrateSteps(db, mockData.Steps); err != nil {
		gut.Error("failed to migrate steps: %w", err)
	}

	if err := migration.MigrateStepAuthors(db, mockData.StepAuthors); err != nil {
		gut.Error("failed to migrate step authors: %w", err)
	}

	if err := migration.MigrateStepEvaluates(db, mockData.StepEvaluates); err != nil {
		gut.Error("failed to migrate step evaluates: %w", err)
	}

	if err := migration.MigrateUserEvaluates(db, mockData.UserEvaluates); err != nil {
		gut.Error("failed to migrate user evaluates: %w", err)
	}

	if err := migration.MigrateStepComments(db, mockData.StepComments); err != nil {
		gut.Error("failed to migrate step comments: %w", err)
	}

	if err := migration.MigrateStepCommentUpvotes(db, mockData.StepCommentUpvotes); err != nil {
		gut.Error("failed to migrate step comment upvotes: %w", err)
	}

	if err := migration.MigrateUserPasses(db, mockData.UserPasses); err != nil {
		gut.Error("failed to migrate user passes: %w", err)
	}

	if err := migration.MigrateEnrolls(db, mockData.Enrolls); err != nil {
		gut.Error("failed to migrate enrolls: %w", err)
	}

	if err := migration.MigrateCourseContents(db, mockData.CourseContents); err != nil {
		gut.Error("failed to migrate course contents: %w", err)
	}

	if err := migration.MigrateUserActivity(db, mockData.UserActivity); err != nil {
		gut.Error("failed to migrate user activity: %w", err)
	}
	
	gut.Debug("Migration completed successfully")
}
