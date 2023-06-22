# Gopenai

OpenAI API wrapper library in Go.  

## Overview

The OpenAI API Library provides a user-friendly interface for interacting with various OpenAI APIs, including but not limited to GPT-3.5 and future releases. It abstracts away the complexities of API communication and authentication, allowing developers to focus on implementing AI-powered features and functionalities without getting bogged down by the technical intricacies.



## Installation
### Prerequsite

``` shell
❯ go version
go version go1.19.5 linux/amd64
```

### Command

```shell 
go get github.com/Point-AI/gopenai
```

## Usage

```Go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Point-AI/gopenai"
)

const API_KEY = "openai-api-key"

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
```

Output:
```shell
As of 2021, the estimated population of the United Kingdom is around 68 million people.
```