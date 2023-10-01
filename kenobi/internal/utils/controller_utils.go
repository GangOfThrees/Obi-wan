package utils

import (
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/models"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/services"
	chatgptservice "github.com/GangOfThrees/Obi-wan/kenobi/internal/services/chatgpt_service"
)

// function to get formatted error response
func GetErrorResponse(reason string, err error) map[string]any {
	if err == nil {
		return map[string]interface{}{
			"message": reason,
			"error":   "An error occured",
		}
	}

	return map[string]interface{}{
		"reason": reason,
		"error":  err,
	}
}

// function to get formatted success response
func GetSuccessResponse(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}

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
