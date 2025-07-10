package main

import (
	"fmt"
	"go_auth/config"
	database "go_auth/databases"
	"go_auth/handlers/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

var cfg = config.MainConfig{}

func main() {
	config.LoadEnv(&cfg)

	dbs := database.DBs{}
	dbs.Init(cfg.DatabaseMaster)

	initController := router.InitController(dbs, cfg)
	app := router.InitRouter(fiber.Config{AppName: cfg.Server.AppName}, initController, cfg)
	if err := app.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)); err != nil {
		log.Fatalln(err)
	}
}
