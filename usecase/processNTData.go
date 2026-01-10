package usecase

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase/fileUC"
	"notion2atlas/usecase/notionUC"
	"notion2atlas/utils"
)

type NDE struct {
	New  []domain.BasePage
	Edit []domain.BasePage
	Del  []string
}

func ProcessNTData[T domain.DBQueryEntity, U domain.BasePage](oldData []T, newData []T, resourceType domain.ResourceType) (*NDE, error) {
	nde, err := GetNDE(oldData, newData)
	if err != nil {
		return nil, err
	}
	fmt.Println("ℹ️ new", len(nde.New))
	fmt.Println("ℹ️ edited", len(nde.Edit))
	fmt.Println("ℹ️ delete", len(nde.Del))
	err = processNDEData[U](*nde, resourceType)
	if err != nil {
		return nil, err
	}
	return nde, nil
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
				if new.GetUpdate() {
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

func toModelIfSlice[T domain.BasePage](src []T) []domain.BasePage {
	result := make([]domain.BasePage, len(src))
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
		err := fileUC.InitCurriculumRelatedDir(id)
		if err != nil {
			fmt.Println("error in usecase/processDelNTData/InitCurriculumRelatedDir")
			return err
		}
		err = fileUC.DelPageByCurriculumId(id)
		if err != nil {
			fmt.Println("error in usecase/processDelNTData/DelPageByCurriculumId")
			return err
		}
		err = fileUC.DelBasePageById(id, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processDelNTData/DelBasePageById")
			return err
		}
	}
	return nil
}

func processEditNTData[T domain.BasePage](editItems []domain.BasePage, resourceType domain.ResourceType) error {
	if len(editItems) <= 0 {
		return nil
	}
	existPageData, err := fileUC.GetPageFile()
	if err != nil {
		fmt.Println("error in usecase/processNTData.go: processEditNTData/fileUC.GetPageFile")
		return err
	}
	var changes []fileUC.ChangeItem
	for _, item := range editItems {
		updatedPages := []domain.NtPageEntity{}
		curriculumPages, err := fileUC.GetPageFile()
		if err != nil {
			fmt.Println("error in usecase/processNTData.go:/processEditNTData/fileUC.GetPageFile")
			return err
		}
		for _, p := range *curriculumPages {
			if p.CurriculumId == item.GetId() {
				ntpage, err := notionUC.GetPageItem(p.Id, resourceType.GetStr())
				if err != nil {
					fmt.Println("error in usecase/processNTData.go: processEditNTData/notinoUC.GetPageItem")
					return err
				}
				isEqualTime, err := ntpage.CompareQueryEntityTime(p)
				if err != nil {
					fmt.Println("error in usecase/processNTData.go: processEditNTData/ntpage.CompareQueryEntityTime")
					return err
				}
				if !isEqualTime {
					updatedPages = append(updatedPages, *ntpage)
				}
			}
		}
		for _, p := range updatedPages {
			err := fileUC.InitPageRelatedFile(p.Id)
			if err != nil {
				fmt.Println("error in usecase/processEditNTData/InitCurriculumRelatedDir")
				return err
			}
			err = saveNtData(p, item.GetId(), *existPageData, resourceType)
			if err != nil {
				fmt.Println("error in usecase/processEditNTData/saveNtData")
				return err
			}
			changeItem, err := fileUC.NewChangeItem(p.Id, p.Title, "update")
			if err != nil {
				fmt.Println("error in usecase/processNTData.go: processEditNTData/fileUC.NewChangeItem")
				return err
			}
			changes = append(changes, *changeItem)
		}
		err = upsertBasePage(item, resourceType)
		if err != nil {
			fmt.Println("error in usecase/saveNtData/UpsertCurriculum")
			return err
		}
	}
	err = fileUC.UpsertChangesFile(changes)
	if err != nil {
		fmt.Println("error in in usecase/processNTData.go: processEditNTData/fileUC/UpsertChangesFile")
		return err
	}
	return nil
}

func processNewNTData[T domain.BasePage](newItems []domain.BasePage, resourceType domain.ResourceType) error {
	if len(newItems) <= 0 {
		return nil
	}
	existPageData, err := fileUC.GetPageFile()
	if err != nil {
		fmt.Println("error in usecase/processNTData.go: processEditNTData/fileUC.GetPageFile")
		return err
	}
	var changes []fileUC.ChangeItem
	for _, item := range newItems {
		ntpage, err := notionUC.GetPageItem(item.GetId(), resourceType.GetStr())
		if err != nil {
			fmt.Println("error in usecase/processNTData.go:/processNweNTData/notionUC.GetPageItem")
			return err
		}
		err = saveNtData(*ntpage, item.GetId(), *existPageData, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processNewNTData/usecase.saveNtData")
			return err
		}
		err = upsertBasePage(item, resourceType)
		if err != nil {
			fmt.Println("error in usecase/saveNtData/UpsertCurriculum")
			return err
		}
		changeItem, err := fileUC.NewChangeItem(item.GetId(), item.GetTitle(), "add")
		if err != nil {
			fmt.Println("error in usecase/processNTData.go: processNewNTData/fileUC.NewChangeItem")
			return err
		}
		changes = append(changes, *changeItem)
	}
	err = fileUC.UpsertChangesFile(changes)
	if err != nil {
		fmt.Println("error in usecase/processNTData.go: processNewNTData/fileUC.UpsertChangesFile")
		return err
	}
	return nil
}

func upsertBasePage(newData domain.BasePage, resourceType domain.ResourceType) error {
	switch resourceType {
	case domain.INFO:
		info, ok := newData.(domain.InfoEntity)
		if !ok {
			fmt.Println("error in usecase/processNTData.go: upsertBasePage/newData.(domain.InfoEntity)")
			return fmt.Errorf("error: failed parse BasePage to InfoEntity")
		}
		err := fileUC.UpsertInfo(info.Id, info, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processNTData.go: upsertBasePage/fileUC.UpsertInfo")
			return err
		}
	case domain.CURRICULUM:
		curr, ok := newData.(domain.CurriculumEntity)
		if !ok {
			fmt.Println("error in usecase/processNTData.go: upsertBasePage/newData.(domain.CurriculumEntity)")
			return fmt.Errorf("error: failed parse BasePage to InfoEntity")
		}
		err := fileUC.UpsertCurriculum(curr.Id, curr, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processNTData.go: upsertBasePage/fileUC.UpsertCurriculum")
			return err
		}
	case domain.ANSWER:
		ans, ok := newData.(domain.AnswerEntity)
		if !ok {
			fmt.Println("error in usecase/processNTData.go: upsertBasePage/newData.(domain.AnswerEntity)")
			return fmt.Errorf("error: failed parse BasePage to InfoEntity")
		}
		err := fileUC.UpsertAnswer(ans.Id, ans, resourceType)
		if err != nil {
			fmt.Println("error in usecase/processNTData.go: upsertBasePage/fileUC.UpsertAnswer")
			return err
		}
	}
	return nil
}
