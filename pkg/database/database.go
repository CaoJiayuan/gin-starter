package database

import (
	"gin-demo/pkg/config"
	"time"

	"github.com/enorith/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DSN        string `yaml:"dsn"`
	LogLevel   string `yaml:"log_level"`
	LogChannel string `yaml:"log_channel"`
	MaxIdle    int    `yaml:"max_idle"`
	MaxOpen    int    `yaml:"max_open"`
}

var conn *gorm.DB

func DB() *gorm.DB {
	return conn
}

func Register(engine *gin.Engine) error {
	var c Config
	config.Unmarshal("database", &c)
	log, err := logging.DefaultManager.Channel(c.LogChannel)
	if err != nil {
		return err
	}

	conf := &gorm.Config{}
	if log != nil {
		conf.Logger = &Logger{
			logLevel:      logger.Info,
			logger:        log,
			SlowThreshold: 300 * time.Millisecond,
		}
	}
	conn, err = gorm.Open(mysql.Open(c.DSN), conf)
	if err == nil {
		db, e := conn.DB()
		if e != nil {
			return e
		}
		db.SetMaxIdleConns(c.MaxIdle)
		db.SetMaxOpenConns(c.MaxOpen)
	}

	return err
}
