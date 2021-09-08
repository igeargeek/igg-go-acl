package main

import (
	"os"

	_ "github.com/igeargeek/igg-go-acl/config"
	"github.com/igeargeek/igg-go-acl/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Server is runing")
	})

	router.V1(app)
	app.Listen(":" + os.Getenv("PORT"))
}
