package util

import (
	"encoding/json"
	"io"
	"net/http"
)

// DecodeByteToJson(resp, &result)
func DecodeResponseToJson(response *http.Response, v any) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return DecodeByteToJson(body, v)
}

// DecodeByteToJson(resp.Content, &result)
func DecodeByteToJson(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
