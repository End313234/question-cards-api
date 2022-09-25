package quizzes

import (
	"question-cards-api/database"
	"question-cards-api/database/models"
	"question-cards-api/utils"

	"github.com/gofiber/fiber/v2"
)

type GetOneQuizQuery struct {
	Id string `json:"id,omitempty"`
}

type GetOneQuizResponse = models.Quiz

func GetOne(ctx *fiber.Ctx) error {
	response := GetOneQuizResponse{}
	query := GetOneQuizQuery{}

	ctx.QueryParser(&query)

	if query.Id == "" {
		return utils.ThrowError(
			ctx,
			fiber.ErrBadRequest.Code,
			"id not found",
			"it looks like you forgot to provide the id of the quiz that you are looking for",
		)
	}

	result := database.Execute(
		"SELECT * FROM quizzes WHERE (id = $id);",
		&response,
		map[string]string{
			"id": "quizzes:" + query.Id,
		},
	)

	if len(result) == 0 {
		return utils.ThrowError(ctx, fiber.ErrNotFound.Code, "quiz not found", "")
	}
	return ctx.Status(fiber.StatusOK).JSON(result[0])
}
