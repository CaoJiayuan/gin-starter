package models

import (
	"gin-demo/pkg/database"

	"github.com/gin-gonic/gin"
)

func Migrator(_ *gin.Engine) error {
	tx := database.DB()

	tx.Migrator().AutoMigrate(&Article{})

	return nil
}
