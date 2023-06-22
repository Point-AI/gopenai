package gopenai

import (
	"context"
	"log"
	"net/http"

	"github.com/Point-AI/gopenai/internal"
)

type CreateEmbeddingsRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
	User  string   `json:"user"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}

type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type CreateEmbeddingsResponse struct {
	Data  []Embedding `json:"data"`
	Usage Usage       `json:"usage"`
}

func (c *Client) CreateEmbeddings(ctx context.Context, req CreateEmbeddingsRequest) (CreateEmbeddingsResponse, error) {
	httpRequest, err := internal.BuildRequest(ctx, BaseURL+embeddingsPathUrl, http.MethodPost, req)
	if err != nil {
		log.Panicf("error while CreateEmbeddings request: %v", err)
		return CreateEmbeddingsResponse{}, err
	}

	response := CreateEmbeddingsResponse{}

	err = c.makeRequest(httpRequest, &response)

	if err != nil {
		return CreateEmbeddingsResponse{}, err
	}

	return response, err
}
