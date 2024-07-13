package models

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/settings"
	"gorm.io/gorm"
)

type (
	Model interface {
	}
	Category struct {
		gorm.Model
		Name        string
		Description string
	}
)

func GetDataset(model interface{}, filters gin.H) {
	if filters["By Key"] != "" {
		settings.Database.Find(model, filters["key"], filters["Key Value"])
	} else {
		settings.Database.Find(model)
	}
}
