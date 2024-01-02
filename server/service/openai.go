package service

import (
	"chat-asisten/config"
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func ResolveComplaint(ctx echo.Context, complaint string) (string, error) {
	config.LoadEnv()

	TOKEN := os.Getenv("OPEN_AI_TOKEN")
	client := openai.NewClient(TOKEN)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Anda adalah seorang asisten virtual yang bertugas untuk membantu pengguna. note: max kata 200",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: complaint,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
