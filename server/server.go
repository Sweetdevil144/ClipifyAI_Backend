package main

import (
	"ClipifyAI/config"
	"ClipifyAI/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitDB()
	app := fiber.New()
	app.Use(cors.New())

	router.Route(app)
	app.Listen(":6969")
}
