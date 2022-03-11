package articles

import (
	"gin-demo/internal/modules/articles/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/api/articles", handlers.PostArticle)
	r.GET("/api/articles", handlers.GetArticles)
	r.DELETE("/api/articles/:id", handlers.DeleteArticles)
}
