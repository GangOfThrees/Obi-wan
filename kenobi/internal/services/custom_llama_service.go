package services

import (
	"os"

	"github.com/Aceix/eli-bot/internal/constants"
	"github.com/Aceix/eli-bot/internal/models"
)

var customLlamaBotServiceInstance *customLlamaBotService

// customLlamaBotService is a struct that represents the chatgpt bot service.
type customLlamaBotService struct {
	ApiKey string
}

type customLlamaBotServiceOption func(*customLlamaBotService)

func NewCustomLlamaBotService(options ...customLlamaBotServiceOption) *customLlamaBotService {
	bot := &customLlamaBotService{
		ApiKey: os.Getenv(constants.CUSTOM_LLAMA_API_KEY),
	}

	for _, option := range options {
		option(bot)
	}

	return bot
}

func (service *customLlamaBotService) InitBotService() error {
	// TODO: send initial prompt
	return nil
}

func (service *customLlamaBotService) GetAnswer(question, conversationId string) (models.BotAnswer, error) {
	// TODO: send question to custom bot api
	return models.BotAnswer{
		BotType: models.CustomLlama,
		Answer:  "i am custom llama",
	}, nil
}

// function for singleton service
func GetCustomLlamaBotServiceInstance() *customLlamaBotService {
	if customLlamaBotServiceInstance == nil {
		customLlamaBotServiceInstance = NewCustomLlamaBotService()
	}
	return customLlamaBotServiceInstance
}
