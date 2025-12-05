package utils

import (
	"fmt"
	"notion2atlas/domain"
)

func IsSameIdInArray[T domain.Entity](targetId string, arr []T) bool {
	for _, item := range arr {
		if item.GetId() == targetId {
			return true
		}
	}
	return false
}

func IsSameIdInMapArray(key string, mapArray []map[string]any, id string) (bool, error) {
	for _, item := range mapArray {
		item_id, ok := item[key]
		if !ok {
			return false, fmt.Errorf("unexpeced data type: item[key]")
		}
		if item_id == id {
			return true, nil
		}
	}
	return false, nil
}
