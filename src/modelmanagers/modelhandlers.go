package modelmanagers

import (
	"github.com/gin-gonic/gin"
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
