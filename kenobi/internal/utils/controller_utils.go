package utils

import (
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/models"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/services"
	chatgptservice "github.com/GangOfThrees/Obi-wan/kenobi/internal/services/chatgpt_service"
)

// function to deduce preferred bot service
func DeduceBotService(preferredService string) models.BotService {
	switch preferredService {
	case "chatgpt":
		return chatgptservice.GetChatGptBotServiceInstance()
	case "custom-llama":
		return services.GetCustomLlamaBotServiceInstance()
	default:
		return chatgptservice.GetChatGptBotServiceInstance()
	}
}
