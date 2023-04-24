package router

import (
	. "strconv"

	"github.com/gin-gonic/gin"
)

func pingEndpoint(c *gin.Context) {
	c.JSON(200, "pong")
}

func InitRoutes(port int) {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/ping", pingEndpoint)
	}

	api.Group("/cdn")
	{
		api.POST("/upload", uploadEndpoint)
		api.GET("/download/:upload_id", downloadEndpoint)
	}

	router.Run(":" + Itoa(port))
}