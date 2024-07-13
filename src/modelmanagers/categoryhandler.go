package modelmanagers

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/models"
	"github.com/goose4Coder/golang-forum-api/settings"
)

func (Manager CategoryManager) Create(values gin.H) gin.H {
	category := models.Category{Name: values["name"].(string), Description: values["description"].(string)}
	settings.Database.Create(&category)
	return gin.H{"Status": "Succesfully created post category named " + values["name"].(string)}
}
func (Manager CategoryManager) Read(filters gin.H) gin.H {
	return gin.H{"Name": "Hello", "Desciption": "Random category"}
}
func (Manager CategoryManager) Update(values gin.H) gin.H {
	return gin.H{}
}

func (Manager CategoryManager) Delete(values gin.H) gin.H {
	return gin.H{}
}
