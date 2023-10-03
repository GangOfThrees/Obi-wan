package controllers

import (
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/repository"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/services"
	lightsabers "github.com/GangOfThrees/Obi-wan/lightsabers/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	var body repository.CreateUserParams

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(lightsabers.GetErrorResponse("Failed to parse request body", err))
	}

	if err := services.ValidateStruct(&body); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(lightsabers.GetErrorResponse("Invalid request body", err))
	}

	newUser, err := services.CreateUserWithCtx(ctx.Context(), body)
	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(lightsabers.GetErrorResponse("Failed to create user", err))
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(lightsabers.GetSuccessResponse("Successfully created user", newUser))
}
