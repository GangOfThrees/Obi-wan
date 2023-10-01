package services

import (
	"os"

	"github.com/GangOfThrees/Obi-wan/internal/constants"
	"github.com/GangOfThrees/Obi-wan/internal/models"
)

var mockBotServiceInstance *mockBotService

// mockBotService is a struct that represents the chatgpt bot service.
type mockBotService struct {
	ApiKey string
}

type mockBotServiceOption func(*mockBotService)

func NewMockBotService(options ...mockBotServiceOption) *mockBotService {
	bot := &mockBotService{
		ApiKey: os.Getenv(constants.MOCK_BOT_API_KEY),
	}

	for _, option := range options {
		option(bot)
	}

	return bot
}

func (service *mockBotService) InitBotService() error {
	// TODO: send initial prompt
	return nil
}

func (service *mockBotService) GetAnswer(question string) (models.BotAnswer, error) {
	// TODO: send question to custom bot api
	return models.BotAnswer{
		BotType: models.ChatGpt,
		Answer:  "42",
	}, nil
}

// function for singleton service
func GetMockBotServiceInstance() *mockBotService {
	if mockBotServiceInstance == nil {
		mockBotServiceInstance = NewMockBotService()
	}
	return mockBotServiceInstance
}
