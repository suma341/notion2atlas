package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/utils"
)

type NDE struct {
	New  []domain.Entity
	Edit []domain.Entity
	Del  []string
}

func ProcessNTData[T domain.DBQueryEntity, U domain.BasePage](oldData []T, newData []T, resourceType domain.ResourceType) error {
	nde, err := GetNDE(oldData, newData)
	if err != nil {
		return err
	}
	fmt.Println("ℹ️ new", len(nde.New))
	fmt.Println("ℹ️ edited", len(nde.Edit))
	fmt.Println("ℹ️ delete", len(nde.Del))
	err = processNDEData[U](*nde, resourceType)
	if err != nil {
		return err
	}
	return nil
}

func GetNDE[T domain.DBQueryEntity](oldData []T, newData []T) (*NDE, error) {
	var deleteList []string
	for _, old := range oldData {
		isExistIdInNewData := utils.IsSameIdInArray(old.GetId(), newData)
		if !isExistIdInNewData {
			deleteList = append(deleteList, old.GetId())
		}
	}
	var newList []T
	for _, new := range newData {
		isNewDataInOldData := utils.IsSameIdInArray(new.GetId(), oldData)
		if !isNewDataInOldData {
			newList = append(newList, new)
		}
	}
	var editedList []T
	for _, new := range newData {
		for _, old := range oldData {
			if new.GetId() == old.GetId() {
				isEqualTime, err := new.CompareQueryEntityTime(old)
				if err != nil {
					fmt.Println("error in usecase/GetNDE/utils.CompareQueryEntityTime")
					return nil, err
				}
				if !isEqualTime && new.GetUpdate() {
					editedList = append(editedList, new)
				}
			}
		}
	}
	return &NDE{
		New:  toModelIfSlice(newList),
		Edit: toModelIfSlice(editedList),
		Del:  deleteList,
	}, nil
}

func toModelIfSlice[T domain.Entity](src []T) []domain.Entity {
	result := make([]domain.Entity, len(src))
	for i, v := range src {
		result[i] = v
	}
	return result
}

func processNDEData[T domain.BasePage](nde NDE, resourceType domain.ResourceType) error {
	err := processNewNTData[T](nde.New, resourceType)
	if err != nil {
		fmt.Println("error in usecase/ProcessNTData/processNewNTData")
		return err
	}
	err = processEditNTData[T](nde.Edit, resourceType)
	if err != nil {
		fmt.Println("error in usecase/processEditNTData/processNewNTData")
		return err
	}
	err = processDelNTData[T](nde.Del, resourceType)
	if err != nil {
		fmt.Println("error in usecase/processDelNTData/processNewNTData")
		return err
	}
	return nil
}

func processDelNTData[T domain.BasePage](delItems []string, resourceType domain.ResourceType) error {
	for _, id := range delItems {
		err := InitCurriculumRelatedDir(id)
		if err != nil {
			fmt.Println("error in usecase/processDelNTData/InitCurriculumRelatedDir")
			return err
		}
		err = DelPageByCurriculumId(id)
		if err != nil {
			fmt.Println("error in usecase/processDelNTData/DelPageByCurriculumId")
			return err
		}
		err = DelBasePageById(id, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processDelNTData/DelBasePageById")
			return err
		}
		err = filemanager.DelFile(constants.OGP_DIR + id + ".png")
		if err != nil {
			fmt.Println("error in usecase/processDelNTData/filemanager.DelFile")
			return err
		}
	}
	return nil
}

func processEditNTData[T domain.BasePage](editItems []domain.Entity, resourceType domain.ResourceType) error {
	for _, item := range editItems {
		err := InitCurriculumRelatedDir(item.GetId())
		if err != nil {
			fmt.Println("error in usecase/processEditNTData/InitCurriculumRelatedDir")
			return err
		}
		err = SaveNotionData[T](item, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processEditNTData/SaveNotionData")
			return err
		}
	}
	return nil
}

func processNewNTData[T domain.BasePage](newItems []domain.Entity, resourceType domain.ResourceType) error {
	for _, item := range newItems {
		err := SaveNotionData[T](item, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processNewNTData/usecase.SaveNotionData")
			return err
		}
	}
	return nil
}
