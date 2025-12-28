package gateway

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/utils"
)

func GetDatFileData[T any](r domain.DatRType) (*T, error) {
	_, tmp := r.GetPath()
	dat, err := filemanager.ReadJson[T](tmp)
	if err != nil {
		fmt.Println("error in fileGW.go/GetDatFileData/ReadJson")
		return nil, err
	}
	return &dat, nil
}

func GetDatPageBlocksFileData(pageId string) ([]domain.AtlBlockEntity, error) {
	path := fmt.Sprintf("%s/%s.dat", constants.PAGE_DATA_DIR, pageId)
	dat, err := filemanager.LoadAndDecodeJson[[]domain.AtlBlockEntity](path)
	if err != nil {
		fmt.Println("error in fileGW.go/GetDatFileData/LoadAndDecodeJson")
		return nil, err
	}
	return *dat, nil
}

func DeleteById(resourceType domain.ResourceType, key string, targetId string) error {
	path, err := resourceType.GetFilePathFromResourceType()
	if err != nil {
		fmt.Println("error in gateway/GetFileData/resourceType.GetFilePathFromResourceType")
		return err
	}
	data, err := filemanager.ReadJSONToMapArray(path)
	if err != nil {
		return fmt.Errorf("error in gateway/DeleteById/filemanager.ReadJSONToMapArray")
	}
	var newList []map[string]any
	for _, c := range data {
		id, err := utils.SafelyRetrieve[string](c, key)
		if err != nil {
			fmt.Println("error in gateway/DeleteById/utils.SafelyRetrieve")
			return err
		}
		if id == nil {
			return fmt.Errorf("unexpeced: id is nil")
		}
		if *id != targetId {
			newList = append(newList, c)
		}
	}
	err = filemanager.WriteJson(newList, path)
	if err != nil {
		fmt.Println("error in gateway/DeleteById/filemanager.ReadJson")
		return err
	}
	return nil
}

func UpsertById[T domain.Entity](resourceType domain.ResourceType, targetId string, newData T) error {
	filepath, err := resourceType.GetFilePathFromResourceType()
	if err != nil {
		fmt.Println("error in gateway/GetFileData/resourceType.GetFilePathFromResourceType")
		return err
	}
	data, err := filemanager.ReadJson[[]T](filepath)
	if err != nil {
		fmt.Println("error in gateway/UpsertById/filemanager.ReadJson")
		return err
	}
	newList := make([]T, 0, len(data))
	found := false
	for _, item := range data {
		if item.GetId() == targetId {
			newList = append(newList, newData)
			found = true
		} else {
			newList = append(newList, item)
		}
	}
	if !found {
		newList = append(newList, newData)
	}

	err = filemanager.WriteJson(newList, filepath)
	if err != nil {
		fmt.Println("error in gateway/UpsertById/filemanager.ReadJson")
		return err
	}
	return nil
}

func UpsertFile[T domain.Entity](resourceType domain.ResourceType, key string, newData []T) ([]T, error) {
	filepath, err := resourceType.GetFilePathFromResourceType()
	if err != nil {
		fmt.Println("error in gateway/GetFileData/resourceType.GetFilePathFromResourceType")
		return nil, err
	}
	arr, err := filemanager.ReadJSONToMapArray(filepath)
	if err != nil {
		fmt.Println("error in gateway/UpsertFile/filemanager.LoadJSONToMap")
		return nil, err
	}
	var newDataMap []map[string]any
	for _, new := range newData {
		map_, err := domain.Struct2Map(new)
		if err != nil {
			fmt.Println("error in gateway/UpsertFile/ domain.Struct2Map")
			return nil, err
		}
		newDataMap = append(newDataMap, map_)
	}
	var newList []map[string]any
	for _, item := range arr {
		id, ok := item[key].(string)
		if !ok {
			return nil, fmt.Errorf("unexpected data type: item[key]")
		}
		isSameIdInNewData, err := utils.IsSameIdInMapArray(key, newDataMap, id)
		if err != nil {
			fmt.Println("error in gateway/UpsertFile/utils.IsSameIdInMapArray")
			return nil, err
		}
		if !isSameIdInNewData {
			newList = append(newList, item)
		}
	}
	newList = append(newList, newDataMap...)
	var newEntities []T
	for _, item := range newList {
		entity, err := domain.Map2Struct[T](item)
		if err != nil {
			fmt.Println("error in gateway/UpsertFile/domain.Map2Struct")
			return nil, err
		}
		if entity == nil {
			return nil, fmt.Errorf("unexpected: entity is nil")
		}
		newEntities = append(newEntities, *entity)
	}
	err = filemanager.WriteJson(newEntities, filepath)
	if err != nil {
		fmt.Println("error in gateway/UpsertFile/filemanager.WriteJson")
		return nil, err
	}
	return newEntities, nil
}
