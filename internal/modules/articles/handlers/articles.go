package handlers

import (
	"gin-demo/internal/models"
	"gin-demo/pkg/database"
	"gin-demo/pkg/response"

	"github.com/gin-gonic/gin"
)

func PostArticle(ctx *gin.Context) {
	var article models.Article
	ctx.Bind(&article)

	database.DB().Save(&article)

	response.Success(nil).End(ctx)
}

func GetArticles(ctx *gin.Context) {
	var articles []models.Article
	tx := database.DB()

	database.PageResponse(tx, &articles, ctx)
}

func DeleteArticles(ctx *gin.Context) {
	tx := database.DB()

	tx.Delete(&models.Article{}, ctx.Param("id"))
}
