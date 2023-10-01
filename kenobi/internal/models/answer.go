package models

type BotAnswer struct {
	BotType SupportedBotService `json:"botType"`
	Answer  string              `json:"answer"`
}
