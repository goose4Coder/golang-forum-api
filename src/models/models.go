package models

import (
	"gorm.io/gorm"
)

type (
	PostBoard struct {
		gorm.Model
		Name        string
		Description string
	}
)
