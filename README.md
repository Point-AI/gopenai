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

	resp, err := client.CreateEmbeddings(context.Background(), gopenai.CreateEmbeddingsRequest{
		Model: gopenai.ModelAdaEmbeddingV2,
		Input: []string{"Hello Wolrd"},
	})

	if err != nil {
		log.Fatalf("error on request: %v", err)
	}

	fmt.Println(resp.Data)

}
```