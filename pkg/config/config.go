package config

import (
	"gin-demo/configs"
	"io/fs"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var AppConfig App

type App struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

func Register(g *gin.Engine) error {
	err := Unmarshal("app", &AppConfig)
	if err != nil {
		return err
	}
	if AppConfig.Port != 0 {
		os.Setenv("PORT", strconv.Itoa(AppConfig.Port))
	}
	if AppConfig.Mode != "" {
		gin.SetMode(AppConfig.Mode)
	}

	return nil
}

func Unmarshal(config string, v interface{}) error {
	data, err := fs.ReadFile(configs.FS, config+".yaml")
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, v)
}
