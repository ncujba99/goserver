package rest

import "github.com/gofiber/fiber/v2"

func Listen() {
	app := fiber.New()

	app.Get("/", index_handeler)

	app.Listen(":3000")
}
