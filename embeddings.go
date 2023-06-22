package gopenai

import (
	"context"
	"log"
	"net/http"

	"github.com/Point-AI/gopenai/internal"
)

const (
	pathUrl = "/embeddings"
)

const (
	ModelAdaSimilarity         = "text-similarity-ada-001"
	ModelBabbageSimilarity     = "text-similarity-babbage-001"
	ModelCurieSimilarity       = "text-similarity-curie-001"
	ModelDavinciSimilarity     = "text-similarity-davinci-001"
	ModelAdaSearchDocument     = "text-search-ada-doc-001"
	ModelAdaSearchQuery        = "text-search-ada-query-001"
	ModelBabbageSearchDocument = "text-search-babbage-doc-001"
	ModelBabbageSearchQuery    = "text-search-babbage-query-001"
	ModelCurieSearchDocument   = "text-search-curie-doc-001"
	ModelCurieSearchQuery      = "text-search-curie-query-001"
	ModelDavinciSearchDocument = "text-search-davinci-doc-001"
	ModelDavinciSearchQuery    = "text-search-davinci-query-001"
	ModelAdaCodeSearchCode     = "code-search-ada-code-001"
	ModelAdaCodeSearchText     = "code-search-ada-text-001"
	ModelBabbageCodeSearchCode = "code-search-babbage-code-001"
	ModelBabbageCodeSearchText = "code-search-babbage-text-001"
	ModelAdaEmbeddingV2        = "text-embedding-ada-002"
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

type CreateEmbeddingsResponse struct {
	Data  []Embedding `json:"data"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

func (c *Client) CreateEmbeddings(ctx context.Context, req CreateEmbeddingsRequest) (CreateEmbeddingsResponse, error) {
	httpRequest, err := internal.BuildRequest(ctx, BaseURL+pathUrl, http.MethodPost, req)
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
