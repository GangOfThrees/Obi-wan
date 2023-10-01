package models

type BotConfig struct {
	BotToken   string
	InitPrompt string
}

type SupportedBotService string

const (
	ChatGpt     SupportedBotService = "chatgpt"
	CustomLlama SupportedBotService = "custom-llama"
)

type BotService interface {
	// function to initialise bot service
	InitBotService() error

	// function to get the answer to a question
	GetAnswer(question, conversationId string) (BotAnswer, error)
}
