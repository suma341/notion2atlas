package filemanager

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func EncodeAndSave(data any, path string) error {
	key := []byte(os.Getenv("ENC_KEY"))
	if len(key) != 32 {
		return errors.New("ENC_KEY must be 32 bytes")
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(jsonBytes); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}
	compressed := buf.Bytes()
	cipherBytes, err := encryptAESGCM(compressed, key)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString(cipherBytes)
	return os.WriteFile(path, []byte(encoded), 0644)
}

func encryptAESGCM(data []byte, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes (AES-256)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}
