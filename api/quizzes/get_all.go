package quizzes

import (
	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("{}")
}
