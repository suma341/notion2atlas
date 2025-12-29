package filemanager

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func LoadAndDecodeJson[T any](path string) (*T, error) {
	// 1. load file (Base64 string)
	encodedBytes, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	// 2. Base64 decode
	cipherBytes, err := base64.StdEncoding.DecodeString(string(encodedBytes))
	if err != nil {
		return nil, err
	}

	// 3. AES decrypt
	key := []byte(os.Getenv("ENC_KEY"))
	if len(key) != 32 {
		return nil, errors.New("ENC_KEY must be 32 bytes")
	}

	decrypted, err := decryptAESGCM(cipherBytes, key)
	if err != nil {
		return nil, err
	}

	// 4. gzip decompress
	gr, err := gzip.NewReader(bytes.NewReader(decrypted))
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	jsonBytes, err := io.ReadAll(gr)
	if err != nil {
		return nil, err
	}

	// 5. JSON unmarshal
	var data T
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func decryptAESGCM(ciphertext []byte, key []byte) ([]byte, error) {
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

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, encrypted := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plain, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, err
	}

	return plain, nil
}

func loadFile(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}
