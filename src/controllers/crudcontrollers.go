package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/modelmanagers"
)

type (
	CRUDController struct {
		Endpoint string
		model    modelmanagers.ModelManager
	}
)

var ApiControllers []CRUDController

func (controller CRUDController) Read(context *gin.Context) {
	context.JSON(200, controller.model.Read(context))
}

func (controller CRUDController) Create(context *gin.Context) {
	context.JSON(200, controller.model.Create(context))
}

func InitApiControllers() {
	ApiControllers = make([]CRUDController, 0, 0)
	ApiControllers = append(ApiControllers, CRUDController{"category", modelmanagers.CategoryManager{}})
}

func RegisterApiControllers(server *gin.Engine) {
	var group *gin.RouterGroup = nil
	for i := 0; i < len(ApiControllers); i++ {
		group = server.Group("/api/v1/" + ApiControllers[i].Endpoint)
		group.GET("/", ApiControllers[i].Read)
		group.POST("/", ApiControllers[i].Create)
	}
}
