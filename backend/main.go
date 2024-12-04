package main

import (
	"backend/internals/config"
	"backend/internals/db"
	"backend/internals/routes"
	"golang.org/x/exp/rand"
	"time"
)

// @title Bookmark API
// @version 1.0
// @description This is baseline server for Bookmark API.
// @contact.name   IoT workshop
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @schemes http https
// @license.name Apache 2.0
// @BasePath /api/v1

func main() {
	rand.Seed(uint64(time.Now().UnixNano()))

	config.BootConfiguration()
	db.SetUpDatabase()
	routes.SetupRoutes()
}
