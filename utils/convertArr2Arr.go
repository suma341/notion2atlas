package utils

import "fmt"

func ConvertArr2Arr[T any, U any](arr []T, convert func(T) (U, error)) ([]U, error) {
	result := make([]U, 0, len(arr))
	for _, item := range arr {
		converted, err := convert(item)
		if err != nil {
			fmt.Println("error in utils/ConvertArr2Arr/convert")
			return nil, err
		}
		result = append(result, converted)
	}
	return result, nil
}
