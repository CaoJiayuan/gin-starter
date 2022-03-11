package main

import (
	"gin-demo/internal"
	"gin-demo/internal/models"
	"gin-demo/pkg/config"
	"gin-demo/pkg/database"
	"gin-demo/pkg/logging"
	"gin-demo/pkg/path"
	"gin-demo/pkg/server"
	"log"
)

func main() {
	s := server.NewServer()
	s.With(config.Register)
	s.With(logging.Regiter(path.BasePath("logs")))
	s.With(database.Register)
	s.With(internal.Routers)
	s.With(models.Migrator)
	e := s.Run()
	if e != nil {
		log.Fatal(e)
	}
}
