package router

import (
	"question-cards-api/api/quizzes"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	router := fiber.New()

	router.Get("/quizzes", quizzes.GetOne)
	router.Get("/quizzes/all", quizzes.GetAll)
	router.Post("/quizzes", quizzes.CreateOne)

	app.Mount("/", router)
}
