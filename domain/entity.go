package domain

import (
	"fmt"
)

type Entity interface {
	GetId() string
}

func EntityIfArr2CategoryArr(ifArr []Entity) ([]CategoryEntity, error) {
	categories, err := ConvertArr2Arr(ifArr, func(i Entity) (CategoryEntity, error) {
		c, ok := i.(CategoryEntity)
		if !ok {
			return CategoryEntity{}, fmt.Errorf("fiailed convert: i.(entity.CategoryEntity)")
		}
		return c, nil
	})
	if err != nil {
		fmt.Println("error in converter/EntityIfArr2CategoryArr/utils.ConvertArr2Arr")
		return nil, err
	}
	return categories, nil
}

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
