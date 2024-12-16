package main

import (
	"bookmark-database/internal/migration"
	"bookmark-database/table"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bsthun/gut"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MockData struct {
	Users              []map[string]interface{} `json:"users"`
	Fields             []map[string]interface{} `json:"fields"`
	Courses            []map[string]interface{} `json:"courses"`
	Modules            []map[string]interface{} `json:"modules"`
	Steps              []map[string]interface{} `json:"steps"`
	StepAuthors        []map[string]interface{} `json:"step_authors"`
	StepEvaluates      []map[string]interface{} `json:"step_evaluates"`
	UserEvaluates      []map[string]interface{} `json:"user_evaluates"`
	StepComments       []map[string]interface{} `json:"step_comments"`
	StepCommentUpvotes []map[string]interface{} `json:"step_comment_upvotes"`
	UserPasses         []map[string]interface{} `json:"user_passes"`
	Enrols             []map[string]interface{} `json:"enrols"`
	CourseContents     []map[string]interface{} `json:"course_contents"`
	Articles           []map[string]interface{} `json:"articles"`
}

func main() {
	// Read and parse mock data
	mockDataBytes, err := os.ReadFile("mockData.json")
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

	dsn := "host=server1.scnn.net user=bookmark password=bookmarkisthebest2024 dbname=bookmark2 port=4040 sslmode=disable"
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

	// Migrate schema
	if err := db.AutoMigrate(
		&table.Article{},
		&table.Course{},
		&table.CourseContent{},
		&table.Enroll{},
		&table.Field{},
		&table.Module{},
		&table.Step{},
		&table.StepAuthor{},
		&table.StepComment{},
		&table.StepCommentUpvote{},
		&table.StepEvaluate{},
		&table.User{},
		&table.UserEvaluate{},
		&table.UserPass{},
	); err != nil {
		gut.Fatal("Failed to migrate schema", err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		gut.Debug("Starting data migration...")

		if err := migration.MigrateUsers(tx, mockData.Users); err != nil {
			return fmt.Errorf("failed to migrate users: %w", err)
		}

		if err := migration.MigrateFields(tx, mockData.Fields); err != nil {
			return fmt.Errorf("failed to migrate fields: %w", err)
		}

		if err := migration.MigrateModules(tx, mockData.Modules); err != nil {
			return fmt.Errorf("failed to migrate modules: %w", err)
		}

		if err := migration.MigrateArticles(tx, mockData.Articles); err != nil {
			return fmt.Errorf("failed to migrate articles: %w", err)
		}

		if err := migration.MigrateCourses(tx, mockData.Courses); err != nil {
			return fmt.Errorf("failed to migrate courses: %w", err)
		}

		if err := migration.MigrateSteps(tx, mockData.Steps); err != nil {
			return fmt.Errorf("failed to migrate steps: %w", err)
		}

		if err := migration.MigrateStepAuthors(tx, mockData.StepAuthors); err != nil {
			return fmt.Errorf("failed to migrate step authors: %w", err)
		}

		if err := migration.MigrateStepEvaluates(tx, mockData.StepEvaluates); err != nil {
			return fmt.Errorf("failed to migrate step evaluates: %w", err)
		}

		if err := migration.MigrateUserEvaluates(tx, mockData.UserEvaluates); err != nil {
			return fmt.Errorf("failed to migrate user evaluates: %w", err)
		}

		if err := migration.MigrateStepComments(tx, mockData.StepComments); err != nil {
			return fmt.Errorf("failed to migrate step comments: %w", err)
		}

		if err := migration.MigrateStepCommentUpvotes(tx, mockData.StepCommentUpvotes); err != nil {
			return fmt.Errorf("failed to migrate step comment upvotes: %w", err)
		}

		if err := migration.MigrateUserPasses(tx, mockData.UserPasses); err != nil {
			return fmt.Errorf("failed to migrate user passes: %w", err)
		}

		if err := migration.MigrateEnrols(tx, mockData.Enrols); err != nil {
			return fmt.Errorf("failed to migrate enrols: %w", err)
		}

		if err := migration.MigrateCourseContents(tx, mockData.CourseContents); err != nil {
			return fmt.Errorf("failed to migrate course contents: %w", err)
		}

		gut.Debug("Data migration completed successfully")
		return nil
	})

	if err != nil {
		gut.Fatal("Migration failed", err)
	}
	gut.Debug("Migration completed successfully")
}