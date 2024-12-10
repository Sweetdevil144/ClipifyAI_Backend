package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handler "ClipifyAI/handler"
)

func Route(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
}