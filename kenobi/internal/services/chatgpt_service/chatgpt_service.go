package chatgptservice

import (
	"errors"
	"os"

	"github.com/Aceix/eli-bot/internal/constants"
	"github.com/Aceix/eli-bot/internal/models"
)

// chatGptBotService is a struct that represents the chatgpt bot service.
type chatGptBotService struct {
	ApiKey string
}

type chatGptBotServiceOption func(*chatGptBotService)

var chatGptBotServiceInstance *chatGptBotService

// TODO: refactor
// customerStore is a map that stores the conversation history of each customer, under each business.
var customerStore = make(map[string][]Message)

// function for singleton service
func GetChatGptBotServiceInstance() *chatGptBotService {
	if chatGptBotServiceInstance == nil {
		chatGptBotServiceInstance = newChatGptBotService()
	}
	return chatGptBotServiceInstance
}

func newChatGptBotService(options ...chatGptBotServiceOption) *chatGptBotService {
	bot := &chatGptBotService{
		ApiKey: os.Getenv(constants.OPEN_AI_API_KEY),
	}
	for _, option := range options {
		option(bot)
	}
	bot.InitBotService()
	return bot
}

func (bot *chatGptBotService) InitBotService() error {
	// TODO: send initial prompt
	return nil
}

func (bot *chatGptBotService) GetAnswer(question, conversationId string) (models.BotAnswer, error) {
	bot.AppendToConversationMessages(
		conversationId,
		Message{
			Role:    "user",
			Content: question,
		},
	)
	msgs, err := bot.GetConversationMessages(conversationId)
	if err != nil {
		return models.BotAnswer{}, err
	}

	res, err := SendChat(ModelGpt35, msgs)
	if err != nil {
		return models.BotAnswer{}, err
	}

	if len(res.Choices) == 0 {
		return models.BotAnswer{}, errors.New("no response answer from bot api")
	}

	ans := res.Choices[0].Message //TODO: verify if this is the correct message
	bot.AppendToConversationMessages(
		conversationId,
		Message{
			Role:    "system",
			Content: ans.Content,
		},
	)

	return models.BotAnswer{
			BotType: models.ChatGpt,
			Answer:  ans.Content,
		},
		nil
}

func (bot *chatGptBotService) GetConversationMessages(conversationId string) ([]Message, error) {
	return customerStore[conversationId], nil
}

func (bot *chatGptBotService) AppendToConversationMessages(conversationId string, msgs ...Message) error {
	customerStore[conversationId] = append(customerStore[conversationId], msgs...)
	return nil
}
