package db

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Gorm *gorm.DB

func SetUpDatabase() {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres",
		*config.Env.DBUsername, *config.Env.DBPassword, *config.Env.DBHost, *config.Env.DBPort)

	// Connect to the Postgres server without specifying a database
	connection := postgres.New(postgres.Config{
		DSN: dbURL,
	})

	db, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
		log.Fatalf("[DATABASE] Unable to connect to Postgres server: %v", err)
	}

	// Check if the target database exists
	dbName := *config.Env.DBName
	var exists bool
	checkQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	if err := db.Raw(checkQuery).Scan(&exists).Error; err != nil {
		log.Fatalf("[DATABASE] Error checking database existence: %v", err)
	}

	// Create the database if it does not exist
	if !exists {
		log.Printf("[DATABASE] Database '%s' does not exist. Creating...", dbName)
		if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)).Error; err != nil {
			log.Fatalf("[DATABASE] Error creating database '%s': %v", dbName, err)
		}
		log.Printf("[DATABASE] Database '%s' created successfully.", dbName)
	}

	// Reconnect to the specific database
	dbURLWithDB := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		*config.Env.DBUsername, *config.Env.DBPassword, *config.Env.DBHost, *config.Env.DBPort, dbName)

	connectionWithDB := postgres.New(postgres.Config{
		DSN: dbURLWithDB,
	})

	Gorm, err = gorm.Open(connectionWithDB, &gorm.Config{})
	if err != nil {
		log.Fatalf("[DATABASE] Unable to connect to database '%s': %v", dbName, err)
	}

	log.Printf("[DATABASE] Connected to database '%s' successfully.", dbName)

	// Initialize model migrations
	if *config.Env.DBAutoMigrate {
		if err := Migrate(); err != nil {
			logrus.Fatal("[Database] Unable to migrate model")
		}
	}
	logrus.Printf("[Database] Initialized postgress connection")
}

func Migrate() error {
	if err := Gorm.AutoMigrate(
		new(models.User),
		new(models.Course),
		new(models.Article),
		new(models.CourseContent),
		new(models.Enroll),
		new(models.FieldType),
		new(models.Module),
		new(models.Step),
		new(models.StepAuthor),
		new(models.StepComment),
		new(models.StepCommentUpvote),
		new(models.StepEvaluate),
		new(models.UserActivity),
		new(models.UserEvaluate),
		new(models.UserPass),
	); err != nil {
		return err
	}
	return nil
}
