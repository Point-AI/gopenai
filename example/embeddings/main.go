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

	resp, err := client.CreateEmbeddings(context.Background(), gopenai.CreateEmbeddingsRequest{
		Model: gopenai.ModelAdaEmbeddingV2,
		Input: []string{"Hello Wolrd"},
	})

	if err != nil {
		log.Fatalf("error on request: %v", err)
	}

	fmt.Println(resp.Data)

}
