package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (model Category) GetData() gin.H {
	toReturn := gin.H{}
	toReturn["Name"] = model.Name
	toReturn["Description"] = model.Description
	toReturn["CreatedAt"] = model.CreatedAt
	toReturn["UpdatedAt"] = model.UpdatedAt
	toReturn["ID"] = model.ID
	return toReturn
}

func (model Category) GetID() string {
	return strconv.Itoa(int(model.ID))
}

func CategoryToDataset(categories []Category) []Model {
	toReturn := make([]Model, 0, 0)
	for i := 0; i < len(categories); i++ {
		toReturn = append(toReturn, categories[i])
	}
	return toReturn
}

// func (model Category) GetAll() Dataset {
// 	toReturn := make([]Category, 0, 0)
// 	settings.Database.Find(&toReturn)
// 	var toRet []Model=toReturn
// }
