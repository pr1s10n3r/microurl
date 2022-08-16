package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).SendString("pong!")
}
