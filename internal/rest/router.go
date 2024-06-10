package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RunAPI() {
	app := fiber.New()

	app.Put("/api/user", createUser)

	fmt.Println("Server listening on port 3000")
	log.Fatal(app.Listen(":3000"))
}
