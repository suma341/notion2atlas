package domain

import (
	"encoding/json"
	"fmt"
)

func Struct2Map[T any](struct_ T) (map[string]any, error) {
	var m map[string]any
	b, err := json.Marshal(struct_)
	if err != nil {
		fmt.Println("error in converter/Struct2Map/json.Marshal(struct_)")
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("error in converter/Struct2Map/json.Unmarshal(b, &m)")
		return nil, err
	}
	return m, nil
}

func Map2Struct[T any](m map[string]any) (*T, error) {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error in converter/Map2Struct/json.Marshal(m)")
		return nil, err
	}

	var s T
	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Println("error in converter/Map2Struct/json.Unmarshal(b, &s)")
		return nil, err
	}

	return &s, nil
}
