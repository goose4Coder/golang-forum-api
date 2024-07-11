package main

import (
	"github.com/goose4Coder/golang-forum-api/models"
	"github.com/goose4Coder/golang-forum-api/settings"
)

func init() {
	settings.ConnectToDatabase()
}

func main() {
	migrateModels()
}

func migrateModels() {
	(*settings.Database).AutoMigrate(&models.PostBoard{})
}
