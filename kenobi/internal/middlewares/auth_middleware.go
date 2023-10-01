package middlewares

import (
	"github.com/Aceix/eli-bot/internal/constants"
	"github.com/Aceix/eli-bot/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func ValidateApiKeyMiddleware(ctx *fiber.Ctx) error {
	allowedKeys := []string{
		"TEST-KEY-1",
		"TEST-KEY-2",
		"TEST-KEY-3",
	}

	for _, key := range allowedKeys {
		if key == ctx.Get(constants.HTTP_HEADER_X_API_KEY) {
			return ctx.Next()
		}
	}

	return ctx.
		Status(fiber.StatusUnauthorized).
		JSON(utils.GetErrorResponse("Invalid API key", constants.ErrInvalidApiKey))
}
