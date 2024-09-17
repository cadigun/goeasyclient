package util

import (
	"bytes"
	"encoding/json"
)

func EncodeObjectToBytesBuffer(data any) (*bytes.Buffer, error) {
	var bytesBuf *bytes.Buffer
	if data != nil {
		payload, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		bytesBuf = bytes.NewBuffer(payload)
	}
	return bytesBuf, nil
}
