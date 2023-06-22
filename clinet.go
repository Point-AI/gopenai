package gopenai

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Point-AI/gopenai/internal"
)

type Client struct {
	authApiToken string

	HTTPClient *http.Client
}

func NewClient(authToken string) *Client {
	return &Client{
		authApiToken: authToken,
		HTTPClient:   &http.Client{},
	}
}

func (c *Client) makeRequest(httpRequest *http.Request, response any) error {
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authApiToken))

	httpResponse, err := c.HTTPClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	if isStatusCodeFailure(httpResponse) {
		return c.handleErrorResp(httpResponse)
	}

	return internal.ParseResponse(httpResponse.Body, response)
}

func isStatusCodeFailure(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil {
		return err
	}

	errRes.Err.HTTPStatusCode = resp.StatusCode

	return errors.New(errRes.Error())
}
