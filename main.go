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
		return c.Redirect("https://forms.gle/Vxd7CrhR8B1j9eEu7", 301)
	})

	app.Listen(":9000")
}
