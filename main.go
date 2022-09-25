package main

import (
	"log"
	"question-cards-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.MountRoutes(app)

	log.Println(app.Listen(":3000"))
}
