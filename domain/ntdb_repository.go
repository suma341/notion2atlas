package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NTDBRepository struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	Title      []struct {
		PlainText string `json:"plain_text"`
	} `json:"title"`
	IsInline bool `json:"is_inline"`
}

func Res2NTDBRepository(res map[string]any) (*NTDBRepository, error) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error in domain/Res2NTDBRepository/son.Marshal(res)")
		return nil, err
	}
	var obj NTDBRepository
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println("error in domain/Res2NTDBRepository/json.Unmarshal(jsonBytes, &obj)")
		return nil, err
	}
	obj.Id = strings.ReplaceAll(obj.Id, "-", "")
	return &obj, nil
}
