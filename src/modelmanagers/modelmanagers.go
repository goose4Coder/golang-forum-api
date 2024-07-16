package modelmanagers

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/models"
)

type (
	ModelManager interface {
		Create(context *gin.Context) gin.H
		Read(context *gin.Context) gin.H
		Update(context *gin.Context) gin.H
		Delete(context *gin.Context) gin.H
	}
	CategoryManager struct {
	}
)

func GetDataset[T models.Model](context *gin.Context) gin.H {
	var body gin.H = gin.H{}
	err := context.BindJSON(&body)
	if body == nil || body["Filter by"] == nil {
		return models.ToData(models.GetAllRows[models.Category]())
	}
	if err != nil {
		body["Status"] = "Error"
		return body
	}
	var data []models.Model
	data, err = models.GetRows[models.Category](body)
	if err != nil {
		body["Status"] = "Error"
		return body
	}
	toReturn := models.ToData(data)
	toReturn["Status"] = "Success"
	return toReturn
}
