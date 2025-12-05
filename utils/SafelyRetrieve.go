package utils

import "fmt"

func SafelyRetrieve[T any](data map[string]any, key string) (*T, error) {
	value, ok := data[key]
	if !ok {
		return nil, fmt.Errorf("no property '%s' in page", key)
	}
	convertedValue, ok := value.(T)
	if !ok {
		return nil, fmt.Errorf("unexpected type: id. can't convert to Type")
	}
	return &convertedValue, nil
}
