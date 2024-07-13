package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/controllers"
)

func StartApiControllers(server *gin.Engine) {
	var group *gin.RouterGroup = nil
	for i := 0; i < len(controllers.ApiControllers); i++ {
		group = server.Group("/api/v1/" + controllers.ApiControllers[i].Endpoint)
		group.GET("/", controllers.ApiControllers[i].Read)
	}
}
