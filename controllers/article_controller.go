package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-project/global"
	"go-project/models"
	"gorm.io/gorm"
	"net/http"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "请求数据有误",
		})
		return
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "系统异常",
		})
		return
	}

	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "系统异常",
		})
		return
	}

	ctx.JSON(http.StatusOK, article)

}

func GetArticles(ctx *gin.Context) {
	var articles []models.Article

	if err := global.Db.Find(&articles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "系统异常",
		})
		return
	}

	ctx.JSON(http.StatusOK, articles)

}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")

	var article models.Article

	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "数据不存在",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "系统异常",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, article)

}
