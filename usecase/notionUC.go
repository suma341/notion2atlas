package usecase

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/gateway"
)

func GetDBQuery(id string) ([]domain.NTDBQueryRepository, error) {
	var data, err = gateway.GetNotionData(domain.DBQuery, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/gateway.GetNotionData")
		return nil, err
	}
	results := data["results"].([]any)
	queryModel, err := domain.Res2NTDBQueryRepository(results)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/domain.Res2NTDBQueryRepository")
		return nil, err
	}
	if queryModel == nil {
		return nil, fmt.Errorf("queryModel is nil")
	}
	return *queryModel, nil
}

func GetChildDB(id string) ([]domain.NTDBQueryRepository, error) {
	var data, err = gateway.GetNotionData(domain.ChildDatabase, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/gateway.GetNotionData")
		return nil, err
	}
	results := data["results"].([]any)
	queryModel, err := domain.Res2NTDBQueryRepository(results)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/domain.Res2NTDBQueryRepository")
		return nil, err
	}
	if queryModel == nil {
		return nil, fmt.Errorf("queryModel is nil")
	}
	return *queryModel, nil
}

func Test(id string) (any, error) {
	var data, err = gateway.GetNotionData(domain.DBQuery, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBQuery/gateway.GetNotionData")
		return nil, err
	}
	// results := data["results"].([]any)
	// queryModel, err := domain.Res2NTDBQueryRepository(results)
	// if err != nil {
	// 	fmt.Println("error in usecase/GetDBQuery/domain.Res2NTDBQueryRepository")
	// 	return nil, err
	// }
	// if queryModel == nil {
	// 	return nil, fmt.Errorf("queryModel is nil")
	// }
	return data, nil
}

func GetDBItem(id string) (*domain.NTDBRepository, error) {
	data, err := gateway.GetNotionData(domain.DB, id)
	if err != nil {
		fmt.Println("error in usecase/GetDBItem/gateway.GetNotionData")
		return nil, err
	}
	filtered, err := domain.Res2NTDBRepository(data)
	if err != nil {
		fmt.Println("error in usecase/GetDBItem/converter.Res2DBModel")
		return nil, err
	}
	return filtered, nil
}

func GetPageItem(id string, type_ string) (*domain.NTPageRepository, error) {
	data, err := gateway.GetNotionData(domain.Page, id)
	if err != nil {
		fmt.Println("error in usecase/GetPageItem/gateway.GetNotionData")
		return nil, err
	}
	filtered, err := domain.ResNTPageRepository(data, type_)
	if err != nil {
		fmt.Println("error in usecase/GetPageItem/domain.ResNTPageRepositoryRepository")
		return nil, err
	}
	return filtered, nil
}
func GetBlockItem(id string) (*domain.NTBlockRepository, error) {
	data, err := gateway.GetNotionData(domain.Block, id)
	if err != nil {
		fmt.Println("error in usecase/GetBlockItem/gateway.GetNotionData")
		return nil, err
	}
	filtered, err := domain.Res2NTBlockRepository(data)
	if err != nil {
		fmt.Println("error in usecase/GetBlockItem/domain.Res2NTBlockRepository")
		return nil, err
	}
	return filtered, nil
}

func GetChildren(id string) ([]domain.NTBlockRepository, error) {
	data, err := gateway.GetNotionData(domain.Children, id)
	if err != nil {
		fmt.Println("error in usecase/GetBlockItem/gateway.GetNotionData")
		return nil, err
	}
	results, ok := data["results"].([]any)
	if !ok {
		return nil, fmt.Errorf("unexpected type: results")
	}
	var list []domain.NTBlockRepository
	for _, item := range results {
		obj := item.(map[string]any)
		blockModel, err := domain.Res2NTBlockRepository(obj)
		if err != nil {
			fmt.Println("error in usecase/GetBlockItem/domain.Res2NTBlockRepository")
			return nil, err
		}
		if blockModel == nil {
			return nil, fmt.Errorf("unexpected: blockModel is nil")
		}
		list = append(list, *blockModel)
	}
	return list, nil
}
