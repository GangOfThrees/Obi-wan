package controllers

import (
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/controllers/dtos"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/services"
	lightsabers "github.com/GangOfThrees/Obi-wan/lightsabers/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CreateBusiness(ctx *fiber.Ctx) error {
	panic("not implemented")

	var body dtos.CreateBusinessDto

	if err := ctx.BodyParser(&body); err != nil {
		log.Error("failed to parse body: %v", err)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(lightsabers.GetErrorResponse("Failed to parse request body", err))
	}

	if err := services.ValidateStruct(&body); err != nil {
		log.Error("failed to validate body: %v", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(lightsabers.GetErrorResponse("Invalid request body", err))
	}

	// services.CreateBusiness()

	return nil
}
