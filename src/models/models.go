package models

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/settings"
	"gorm.io/gorm"
)

type (
	Model interface {
		GetData() gin.H
		GetID() string
	}
	Category struct {
		gorm.Model
		Name        string
		Description string
	}
)

const ITEMS_PER_REQUEST int = 60

func GetAllRows[T Model]() []Model {
	var model []T = make([]T, 0, 0)
	settings.Database.Find(&model)
	toReturn := make([]Model, 0, 0)
	for i := 0; i < len(model) && i < ITEMS_PER_REQUEST; i++ {
		toReturn = append(toReturn, model[i])
	}
	return toReturn
}

func GetRows[T Model](filters gin.H) ([]Model, error) {
	var model []T = make([]T, 0, 0)
	err := settings.Database.Find(&model, filters["Filter by"], filters["Value"]).Error
	if err != nil {
		return []Model{}, err
	}
	var count int64 = 0
	count = int64(filters["Count"].(float64))
	if err != nil {
		filters["Count"] = ITEMS_PER_REQUEST
	}
	toReturn := make([]Model, 0, 0)
	for i := 0; i < len(model) && i < int(count); i++ {
		toReturn = append(toReturn, model[i])
	}
	return toReturn, nil
}

func ToData(dataset []Model) gin.H {
	toReturn := gin.H{}
	for i := 0; i < len(dataset); i++ {
		toReturn[dataset[i].GetID()] = dataset[i].GetData()
	}
	return toReturn
}
