package main

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAi struct {
	client *openai.Client
}

func NewOpenAi() *OpenAi {
	return &OpenAi{
		client: openai.NewClient(cfg.OpenKey),
	}
}

func (o *OpenAi) AskForOpenAI(ctx context.Context, user, text string) (string, error) {
	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
