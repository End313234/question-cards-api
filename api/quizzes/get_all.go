package quizzes

import (
	"question-cards-api/database"
	"question-cards-api/database/models"

	"github.com/gofiber/fiber/v2"
)

type GetAllResponse = models.Quiz

func GetAll(ctx *fiber.Ctx) error {
	response := GetAllResponse{}
	result := database.Execute("SELECT * FROM quizzes;", &response, map[string]string{})
	return ctx.Status(fiber.StatusOK).JSON(result)
}
