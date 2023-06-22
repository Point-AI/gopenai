package internal

import (
	"encoding/json"
	"io"
)

func ParseResponse(body io.Reader, response any) error {
	if response == nil {
		return nil
	}

	if output, ok := response.(*string); ok {
		parseString(body, output)
	}

	return json.NewDecoder(body).Decode(response)
}

func parseString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	*output = string(b)
	return nil
}
