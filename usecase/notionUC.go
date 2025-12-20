package usecase

import (
	"errors"
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/gateway"
)

func GetDBQuery(id string) ([]domain.NtDBQueryEntity, error) {
	var data, err = gateway.GetNotionData(domain.DBQuery, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/gateway.GetNotionData")
		return nil, err
	}
	results := data["results"].([]any)
	queryModel, err := domain.Res2NtDBQueryEntity(results)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/domain.Res2NTDBQueryEntity")
		return nil, err
	}
	if queryModel == nil {
		return nil, fmt.Errorf("queryModel is nil")
	}
	return *queryModel, nil
}

func GetChildDB(id string) ([]domain.NtDBQueryEntity, error) {
	var data, err = gateway.GetNotionData(domain.ChildDatabase, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/gateway.GetNotionData")
		return nil, err
	}
	results := data["results"].([]any)
	queryModel, err := domain.Res2NtDBQueryEntity(results)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/domain.Res2NTDBQueryEntity")
		return nil, err
	}
	if queryModel == nil {
		return nil, fmt.Errorf("queryModel is nil")
	}
	return *queryModel, nil
}

func Test(id string) (any, error) {
	var data, err = gateway.GetNotionData(domain.Block, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/gateway.GetNotionData")
		return nil, err
	}
	// results := data["results"].([]any)
	// queryModel, err := domain.Res2NTDBQueryEntity(results)
	// if err != nil {
	// 	fmt.Println("error in usecase/GetDBQuery/domain.Res2NTDBQueryEntity")
	// 	return nil, err
	// }
	// if queryModel == nil {
	// 	return nil, fmt.Errorf("queryModel is nil")
	// }
	return data, nil
}

func GetDBItem(id string) (*domain.NtDBEntity, error) {
	data, err := gateway.GetNotionData(domain.DB, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBItem/gateway.GetNotionData")
		return nil, err
	}
	filtered, err := domain.Res2NtDBEntity(data)
	if err != nil {
		if errors.Is(err, domain.ErrNotionErrorResponse) {
			return nil, domain.ErrNotionErrorResponse
		}
		fmt.Println("error in usecase/GetDBItem/converter.Res2DBModel")
		return nil, err
	}
	return filtered, nil
}

func GetPageItem(id string, type_ string) (*domain.NtPageEntity, error) {
	data, err := gateway.GetNotionData(domain.Page, id)
	if err != nil {
		fmt.Println("error in usecase/GetPageItem/gateway.GetNotionData")
		return nil, err
	}
	filtered, err := domain.ResNtPageEntity(data, type_)
	if err != nil {
		if errors.Is(err, domain.ErrNotionErrorResponse) {
			return nil, domain.ErrNotionErrorResponse
		}
		fmt.Println("error in usecase/GetPageItem/domain.ResNtPageEntityEntity")
		return nil, err
	}
	return filtered, nil
}

func GetBlockItem(id string) (*domain.NTBlockEntity, error) {
	data, err := gateway.GetNotionData(domain.Block, id)
	if err != nil {
		fmt.Println("error in usecase/GetBlockItem/gateway.GetNotionData")
		return nil, err
	}
	filtered, err := domain.Res2NTBlockEntity(data)
	if err != nil {
		if errors.Is(err, domain.ErrNotionErrorResponse) {
			return nil, domain.ErrNotionErrorResponse
		}
		fmt.Println("error in usecase/GetBlockItem/domain.Res2NTBlockEntity")
		return nil, err
	}
	return filtered, nil
}

func GetChildren(id string) ([]domain.NTBlockEntity, error) {
	data, err := gateway.GetNotionData(domain.Children, id)
	if err != nil {
		fmt.Println("error in usecase/GetBlockItem/gateway.GetNotionData")
		return nil, err
	}
	results, ok := data["results"].([]any)
	if !ok {
		return nil, fmt.Errorf("unexpected type: results")
	}
	var list []domain.NTBlockEntity
	for _, item := range results {
		obj := item.(map[string]any)
		blockModel, err := domain.Res2NTBlockEntity(obj)
		if err != nil {
			if !errors.Is(err, domain.ErrNotionErrorResponse) {
				fmt.Println("error in usecase/GetBlockItem/domain.Res2NTBlockEntity")
				return nil, err
			}
		}
		if blockModel != nil {
			list = append(list, *blockModel)
		}
	}
	return list, nil
}
