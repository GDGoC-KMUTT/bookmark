package db

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func SetUpDatabase() {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Env.DBUsername, config.Env.DBPassword, config.Env.DBHost, config.Env.DBPort, config.Env.DBName)

	// open Postgres connection
	connection := postgres.New(postgres.Config{
		DSN: dbURL,
	})

	if db, err := gorm.Open(connection, &gorm.Config{}); err != nil {
		logrus.Fatal("[DATABASE] Unable to load postgres database")
	} else {
		Gorm = db
	}

	// Initialize model migrations
	if config.Env.DBAutoMigrate {
		if err := Migrate(); err != nil {
			logrus.Fatal("[Database] Unable to migrate model")
		}
	}
	AssignModel()
	logrus.Debug("[Database] Initialized postgress connection")
}

func Migrate() error {
	if err := Gorm.AutoMigrate(
		new(models.User),
	); err != nil {
		return err
	}
	return nil
}
