package router

import (
	"fmt"
	"go_auth/config"
	"go_auth/handlers/controllers/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"go.elastic.co/apm/module/apmfiber"
)

func InitRouter(config fiber.Config, ctrl controllerCollections, cfg config.MainConfig) *fiber.App {
	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(apmfiber.Middleware())
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("uri_path", fmt.Sprintf("%v:%v", c.Method(), c.Path()))
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Welcome")
	})

	app.Get("/users", middlewares.JWTMiddleware(), ctrl.userController.ListUsers)
	app.Post("/register", ctrl.userController.Register)
	app.Post("/login", ctrl.userController.Login)

	return app
}
