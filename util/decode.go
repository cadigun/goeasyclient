package util

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
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

func DecodeFromFileToByte(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func DecodeFromFileToType(filepath string, v any) error {
	data, err := DecodeFromFileToByte(filepath)
	if err != nil {
		return err
	}
	return DecodeByteToJson(data, v)
}
