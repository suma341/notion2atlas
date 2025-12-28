package filemanager

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
)

func LoadAndDecodeJson[T any](path string) (*T, error) {
	encoded, err := LoadFile(path)
	if err != nil {
		return nil, err
	}

	jsonStr, err := GzipBase64Decode(encoded)
	if err != nil {
		return nil, err
	}

	return UnmarshalJsonStr[T](jsonStr)
}

func UnmarshalJsonStr[T any](jsonStr string) (*T, error) {
	var data T
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func LoadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GzipBase64Decode(encoded string) (string, error) {
	compressed, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	gr, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return "", err
	}
	defer gr.Close()

	decodedBytes, err := io.ReadAll(gr)
	if err != nil {
		return "", err
	}

	return string(decodedBytes), nil
}
