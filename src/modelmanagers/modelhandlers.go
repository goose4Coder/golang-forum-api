package modelmanagers

import (
	"github.com/gin-gonic/gin"
)

type (
	ModelManager interface {
		Create(values gin.H) gin.H
		Read(filters gin.H) gin.H
		Update(values gin.H) gin.H
		Delete(values gin.H) gin.H
	}
	CategoryManager struct {
	}
)
