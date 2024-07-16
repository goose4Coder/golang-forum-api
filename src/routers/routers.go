package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/controllers"
)

func RegisterApiControllers(server *gin.Engine) {
	var group *gin.RouterGroup = nil
	for i := 0; i < len(controllers.ApiControllers); i++ {
		group = server.Group("/api/v1/" + controllers.ApiControllers[i].Endpoint)
		group.POST("/", controllers.ApiControllers[i].Create)
		group.GET("/", controllers.ApiControllers[i].Read)
		group.PUT("/:id/", controllers.ApiControllers[i].Update)
	}
}
