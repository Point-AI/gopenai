package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Point-AI/gopenai"
)

const API_KEY = "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"

func main() {
	client := gopenai.NewClient(API_KEY)

	response, err := client.CreateChatCompletion(context.Background(), gopenai.ChatCompletionRequest{
		Model: gopenai.ModelGPT3Dot5Turbo,
		Messages: []gopenai.ChatCompletionMessage{
			{
				Role:    gopenai.ChatMessageRoleSystem,
				Content: "You are intelligent assistant!",
			},
			{
				Role:    gopenai.ChatMessageRoleUser,
				Content: "What is the capital of great britain?",
			},
			{
				Role:    gopenai.ChatMessageRoleAssistant,
				Content: "London is capital of Great Britain",
			},
			{
				Role:    gopenai.ChatMessageRoleUser,
				Content: "What is the population of that country?",
			},
		},
		Temperature: 0.5,
	})
	if err != nil {
		log.Fatalf("error while http request: %v", err)
	}

	fmt.Println(response.Choices[0].Message.Content)
}
