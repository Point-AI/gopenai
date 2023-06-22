package gopenai

import "fmt"

// OpenAI error response handling
type ErrorResponse struct {
	Err *OpenAPIError `json:"error"`
}

type OpenAPIError struct {
	Message        string  `json:"message"`
	Type           string  `json:"type"`
	Param          *string `json:"param,omitempty"`
	Code           any     `json:"code,omitempty"`
	HTTPStatusCode int     `json:"-"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("http status code: %d, error message: %s", e.Err.HTTPStatusCode, e.Err.Message)
}
