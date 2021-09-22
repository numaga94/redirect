package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cache.New())

	app.Get("/ski-tignes-2021", func(c *fiber.Ctx) error {
		return c.Redirect("https://docs.google.com/forms/d/e/1FAIpQLSdK43WI3EH-9WinmF_CRRe8LHCmwuuH2yuGvYyCvwMEo1Ed0w/viewform", 301)
	})

	app.Listen(":8005")
}
