package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NtDBEntity struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	Title      []struct {
		PlainText string `json:"plain_text"`
	} `json:"title"`
	IsInline bool   `json:"is_inline"`
	Object   string `json:"object"` // database
}

func Res2NtDBEntity(res map[string]any) (*NtDBEntity, error) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error in domain/Res2NtDBEntity/son.Marshal(res)")
		return nil, err
	}
	var obj NtObjectProbe
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println("error in domain/Res2NTBlockEntity/son.Unmarshal(jsonBytes, &obj)")
		return nil, err
	}
	switch obj.Object {
	case "database":
		var block NtDBEntity
		if err := json.Unmarshal(jsonBytes, &block); err != nil {
			return nil, err
		}
		block.Id = strings.ReplaceAll(block.Id, "-", "")
		return &block, nil
	case "error":
		var ntErr NTErrorEntity
		if err := json.Unmarshal(jsonBytes, &ntErr); err != nil {
			return nil, err
		}
		return nil, ErrNotionErrorResponse
	}
	return nil, fmt.Errorf("unexpected: object type is Neither database or error")
}
