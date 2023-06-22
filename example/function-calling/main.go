package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Point-AI/gopenai"
)

const API_KEY = "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"

type WeatherInfo struct {
	Location    string   `json:"location"`
	Temperature int      `json:"temperature"`
	Unit        string   `json:"unit"`
	Forecast    []string `json:"forecast"`
}
type Args struct {
	Location string `json:"location"`
	Unit     string `json:"unit"`
}

func GetCurrentWeather(location string, unit string) WeatherInfo {

	// Getting weather information for the internet

	response := WeatherInfo{
		Location:    location,
		Temperature: 30,
		Unit:        unit,
		Forecast:    []string{"sunny", "windy"},
	}

	return response
}

func main() {
	clinet := gopenai.NewClient(API_KEY)

	functions := []gopenai.FunctionDefinition{
		{
			Name:        "GetCurrentWeather",
			Description: "Get the current weather",
			Parameters: gopenai.JSONSchemaDefinition{
				Type: gopenai.JSONSchemaObject,
				Properties: map[string]gopenai.JSONSchemaDefinition{
					"location": {
						Type:        "string",
						Description: "The city and state, e.g. San Francisco, CA",
					},
					"unit": {
						Type:        "string",
						Description: "The temperature unit to use. Infer this from the users location.",
						Enum:        []string{"celsius", "fahrenheit"},
					},
				},
				Required: []string{"location", "unit"},
			},
		},
	}

	response, err := clinet.CreateChatCompletion(context.Background(), gopenai.ChatCompletionRequest{
		Model: gopenai.ModelGPT3Dot5Turbo0613,
		Messages: []gopenai.ChatCompletionMessage{
			{
				Role:    gopenai.ChatMessageRoleSystem,
				Content: "Don't make assumptions about what values to plug into functions. Ask for clarification if a user request is ambiguous",
			},
			{
				Role:    gopenai.ChatMessageRoleUser,
				Content: "What's the weather like today ?",
			},
			{
				Role:    gopenai.ChatMessageRoleAssistant,
				Content: "In which city and state would you like to know the current weather?",
			},
			{
				Role:    gopenai.ChatMessageRoleUser,
				Content: "I live in Tashkent, Uzbekistan.",
			},
		},
		Functions: functions,
	})
	if err != nil {
		log.Fatalf("error while http request: %v", err)
	}

	if response.Choices[0].FinishReason == gopenai.FinishReasonFunctionCall {
		fucntionToCall := response.Choices[0].Message.FunctionCall

		if fucntionToCall.Name == "GetCurrentWeather" {
			var args Args
			if err = json.Unmarshal([]byte(fucntionToCall.Arguments), &args); err != nil {
				log.Fatalf("error while parsing arguments: %v", err)
			}

			fmt.Println(GetCurrentWeather(args.Location, args.Unit))
		}
	}
}
