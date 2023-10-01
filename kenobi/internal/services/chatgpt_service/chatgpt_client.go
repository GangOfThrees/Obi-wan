package chatgptservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Aceix/eli-bot/internal/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetChatGptBaseUrl() string {
	return os.Getenv(constants.CHAT_GPT_BASE_URL)
}

func SendChat(model string, messages []Message) (CompletionObject, error) {
	res := CompletionObject{}

	agent := fiber.Post(GetChatGptBaseUrl())
	reqBody := CompletionReqDto{
		Model:       model,
		Messages:    messages,
		MaxTokens:   200,
		Temperature: 0.2,
	}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return res, err
	}
	log.Infof("%s", b)
	agent.Set("Authorization", fmt.Sprintf("Bearer %s", GetChatGptBotServiceInstance().ApiKey))
	agent.Set("Content-Type", "application/json")
	agent.Body(b)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return res, errors.Join(errs...)
	}
	if statusCode/200 != 1 {
		return res, fmt.Errorf("unsuccessful response from OpenAI: %s", string(body))
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
