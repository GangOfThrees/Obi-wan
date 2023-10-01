package controllers

import (
	"github.com/Aceix/eli-bot/internal/constants"
	"github.com/Aceix/eli-bot/internal/controllers/dtos"
	"github.com/Aceix/eli-bot/internal/services"
	"github.com/Aceix/eli-bot/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// KnowledgeBase is a function that returns a response to questions asked.
func KnowledgeBase(ctx *fiber.Ctx) error {
	var body dtos.QuestionDto
	if err := ctx.BodyParser(&body); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(utils.GetErrorResponse("Failed to parse request body", err))
	}

	err := services.ValidateStruct(&body)
	if err != nil {
		log.Errorf("Failed to validate body: %v", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(utils.GetErrorResponse("Invalid request body", err))
	}

	reqHeaders := ctx.GetReqHeaders()
	preferredBot := utils.DeduceBotService(reqHeaders[constants.HTTP_HEADER_X_BOT_SERVICE])

	session, err := ConvoSessionStore.Get(ctx)
	if err != nil {
		log.Errorf("Failed to get convo session: %v", err)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(utils.GetErrorResponse("Failed to get convo session", err))
	}

	chatSessionId := session.ID()

	ans, err := preferredBot.GetAnswer(body.Question, chatSessionId)
	if err != nil {
		log.Errorf("Failed to get answer: %v", err)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(utils.GetErrorResponse("Failed to get answer", err))
	}

	if err = session.Save(); err != nil {
		log.Errorf("Failed to save session: %s", err)
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(utils.GetSuccessResponse("Successfully retrieved answer", dtos.ToAnswerDto(ans)))
}