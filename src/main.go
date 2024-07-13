package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goose4Coder/golang-forum-api/controllers"
	"github.com/goose4Coder/golang-forum-api/settings"
)

func init() {
	settings.ConnectToDatabase()
	controllers.InitApiControllers()
}

func main() {
	server := gin.Default()
	controllers.RegisterApiControllers(server)
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})
	server.Run()
}
