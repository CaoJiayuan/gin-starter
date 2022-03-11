package internal

import (
	"gin-demo/internal/modules/articles"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) error {
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	articles.Routes(r)
	return nil
}
