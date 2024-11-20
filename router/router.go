package router

import (
	"github.com/gin-gonic/gin"
	"go-project/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API 分组
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Registry)
	}

	return r
}
