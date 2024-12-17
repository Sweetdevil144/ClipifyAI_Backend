package router

import (
	"ClipifyAI/handler"
	users "ClipifyAI/handler/users"
	"ClipifyAI/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Route(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	user := api.Group("/user")
	user.Get("/register", users.CreateUser)
	user.Get("/login", users.LoginUser)
	user.Use(middleware.Protected())
	user.Get("/updateApiKeys", users.UpdateKeys)
}
