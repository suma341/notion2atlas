package filemanager

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"os"
)

func EncodeAndSave(data any, path string) error {
	jsonStr := jsonToString(data)
	encoded, err := gzipBase64Encode(jsonStr)
	if err != nil {
		return err
	}
	err = saveEncodedToFile(path, encoded)
	if err != nil {
		return err
	}
	return nil
}

func jsonToString(data any) string {
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}

func gzipBase64Encode(input string) (string, error) {
	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(input))
	if err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encoded, nil
}

func saveEncodedToFile(path string, encoded string) error {
	return os.WriteFile(path, []byte(encoded), 0644)
}
