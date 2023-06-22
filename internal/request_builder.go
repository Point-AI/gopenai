package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func BuildRequest(ctx context.Context, url string, method string, body any) (*http.Request, error) {
	if body == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	return request, nil
}


