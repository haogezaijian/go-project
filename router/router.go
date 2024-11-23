package router

import (
	"github.com/gin-gonic/gin"
	"go-project/controllers"
	"go-project/middlewares"
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

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticleById)

		api.POST("/article/:id/like", controllers.LikeArticle)
		api.GET("/article/:id/like", controllers.GetArticleLikes)

	}
	return r
}
