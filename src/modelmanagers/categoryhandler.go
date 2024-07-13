package modelmanagers

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/models"
	"github.com/goose4Coder/golang-forum-api/settings"
)

func (Manager CategoryManager) Create(context *gin.Context) gin.H {
	var body gin.H
	context.Bind(&body)
	if body["Name"].(string) != "" {
		category := models.Category{Name: body["Name"].(string), Description: body["Description"].(string)}
		settings.Database.Create(&category)
		return gin.H{"Status": "Succesfully created post category named " + body["Name"].(string)}
	}
	return gin.H{"Status": "Failed to create a post category"}
}
func (Manager CategoryManager) Read(context *gin.Context) gin.H {
	var category models.Category
	var body gin.H
	context.Bind(&body)
	models.GetDataset(&category, body)
	return gin.H{"Name": category.Name, "Description": category.Description}
}
func (Manager CategoryManager) Update(context *gin.Context) gin.H {
	return gin.H{}
}

func (Manager CategoryManager) Delete(context *gin.Context) gin.H {
	return gin.H{}
}
