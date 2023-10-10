package controllers

import (
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/constants"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/controllers/dtos"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/services"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/utils"
	lightsabers "github.com/GangOfThrees/Obi-wan/lightsabers/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// KnowledgeBase is a function that returns a response to questions asked.
func KnowledgeBase(ctx *fiber.Ctx) error {
	var body dtos.QuestionDto
	if err := ctx.BodyParser(&body); err != nil {
		log.Errorf("failed to parse request body: ", err)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(lightsabers.GetErrorResponse("Failed to parse request body", err))
	}

	err := services.ValidateStruct(&body)
	if err != nil {
		log.Errorf("failed to validate body: ", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(lightsabers.GetErrorResponse("Invalid request body", err))
	}

	reqHeaders := ctx.GetReqHeaders()
	preferredBot := utils.DeduceBotService(reqHeaders[constants.HTTP_HEADER_X_BOT_SERVICE])

	session, err := ConvoSessionStore.Get(ctx)
	if err != nil {
		log.Errorf("failed to get convo session: ", err)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(lightsabers.GetErrorResponse("Failed to get convo session", err))
	}

	chatSessionId := session.ID()

	ans, err := preferredBot.GetAnswer(body.Question, chatSessionId)
	if err != nil {
		log.Errorf("failed to get answer: ", err)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(lightsabers.GetErrorResponse("Failed to get answer", err))
	}

	if err = session.Save(); err != nil {
		log.Errorf("failed to save session: ", err)
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(lightsabers.GetSuccessResponse("Successfully retrieved answer", dtos.ToAnswerDto(ans)))
}
