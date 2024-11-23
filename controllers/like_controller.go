package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go-project/global"
	"net/http"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId + ":likes"

	if err := global.RedisDb.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "系统异常",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "点赞成功",
	})
}

func GetArticleLikes(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId + ":likes"

	count, err := global.RedisDb.Get(likeKey).Result()

	if err == redis.Nil {
		count = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "系统异常",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
