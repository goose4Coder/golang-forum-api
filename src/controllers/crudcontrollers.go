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

func (controller CRUDController) Create(context *gin.Context) {
	context.JSON(200, controller.model.Create(context))
}

func (controller CRUDController) Read(context *gin.Context) {
	context.JSON(200, controller.model.Read(context))
}

func (controller CRUDController) Update(context *gin.Context) {
	context.JSON(200, controller.model.Update(context))
}

func InitApiControllers() {
	ApiControllers = make([]CRUDController, 0, 0)
	ApiControllers = append(ApiControllers, CRUDController{"category", modelmanagers.CategoryManager{}})
}
