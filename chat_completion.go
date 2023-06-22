package gopenai

import (
	"context"
	"net/http"

	"github.com/Point-AI/gopenai/internal"
)

// *****************
// REQUEST MODEL for chat competion
// *****************

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleFunction  = "function"
)

type FunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

type ChatCompletionMessage struct {
	Role         string        `json:"role"`
	Content      string        `json:"content"`
	Name         string        `json:"name,omitempty"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

type FunctionDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// The JSONSchemaDefinition struct, or byte arrat can be passed
	Parameters any `json:"parameters"`
}

// Request model for chat completion
type ChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	MaxTokens        int                     `json:"max_tokens,omitempty"`
	Temperature      float32                 `json:"temperature,omitempty"`
	TopP             float32                 `json:"top_p,omitempty"`
	N                int                     `json:"n,omitempty"`
	Stream           bool                    `json:"stream,omitempty"`
	Stop             []string                `json:"stop,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                 `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int          `json:"logit_bias,omitempty"`
	User             string                  `json:"user,omitempty"`
	Functions        []FunctionDefinition    `json:"functions,omitempty"`
	FunctionCall     any                     `json:"function_call,omitempty"`
}

// *****************
// RESPONSE MODEL for chat competion
// *****************

const (
	FinishReasonStop          = "stop"           // stop: API returned complete message, or a message terminated by one of the stop sequences provided via the stop parameter
	FinishReasonLength        = "length"         // length: Incomplete model output due to max_tokens parameter or token limit
	FinishReasonFunctionCall  = "function_call"  // function_call: The model decided to call a function
	FinishReasonContentFilter = "content_filter" // content_filter: Omitted content due to a flag from our content filters
	FinishReasonNull          = "null"           // null: API response still in progress or incomplete
)

type ChatCompletionChoice struct {
	Index        int                   `json:"index"`
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
}

// Response model for chat completion
type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

func (c *Client) CreateChatCompletion(ctx context.Context, req ChatCompletionRequest) (ChatCompletionResponse, error) {
	httpRequest, err := internal.BuildRequest(ctx, BaseURL+chatCompletionPathUrl, http.MethodPost, req)
	if err != nil {
		return ChatCompletionResponse{}, err
	}

	resp := ChatCompletionResponse{}

	err = c.makeRequest(httpRequest, &resp)
	if err != nil {
		return ChatCompletionResponse{}, err
	}

	return resp, nil
}
