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
	encoded, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	jsonStr, err := gzipBase64Decode(encoded)
	if err != nil {
		return nil, err
	}

	return unmarshalJsonStr[T](jsonStr)
}

func unmarshalJsonStr[T any](jsonStr string) (*T, error) {
	var data T
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func loadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func gzipBase64Decode(encoded string) (string, error) {
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
