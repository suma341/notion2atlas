package fileUC

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
)

func GetPageDataFile(pageId string) ([]domain.AtlBlockEntity, error) {
	data, err := gateway.GetDatPageBlocksFileData(pageId)
	if err != nil {
		fmt.Println("error in usecase/GetPageDataFile/gateway.GetDatPageBlocksFileData")
		return nil, err
	}
	return data, nil
}

func GetPageFile() (*[]domain.AtlPageEntity, error) {
	data, err := gateway.GetDatFileData[[]domain.AtlPageEntity](domain.PAGE_DAT)
	if err != nil {
		fmt.Println("error in usecase/GetPageFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func GetCurriculumFile() (*[]domain.CurriculumEntity, error) {
	data, err := gateway.GetDatFileData[[]domain.CurriculumEntity](domain.CURRICULUM_DAT)
	if err != nil {
		fmt.Println("error in usecase/GetCurriculumFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func GetCategoryFile() (*[]domain.CategoryEntity, error) {
	data, err := gateway.GetDatFileData[[]domain.CategoryEntity](domain.CATEGORY_DAT)
	if err != nil {
		fmt.Println("error in usecase/GetCurriculumFile/gateway.GetDatFileData")
		return nil, err
	}
	return data, nil
}

func GetInfoFile() (*[]domain.InfoEntity, error) {
	data, err := gateway.GetDatFileData[[]domain.InfoEntity](domain.INFO_DAT)
	if err != nil {
		fmt.Println("error in usecase/GetInfoFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func GetAnswerFile() (*[]domain.AnswerEntity, error) {
	data, err := gateway.GetDatFileData[[]domain.AnswerEntity](domain.ANSWER_DAT)
	if err != nil {
		fmt.Println("error in usecase/GetAnswerFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func DelBasePageById(id string, resourseType domain.ResourceType) error {
	err := gateway.DeleteById(resourseType, "id", id)
	if err != nil {
		fmt.Println("error in usecase/DelCurriculumById/gateway.DeleteById")
		return err
	}
	return nil
}

func DelPageByCurriculumId(curriculumId string) error {
	err := gateway.DeleteById(domain.PAGE, "curriculumId", curriculumId)
	if err != nil {
		fmt.Println("error in usecase/DelPageById/gateway.DeleteById")
		return err
	}
	return nil
}

func DelCategoryById(categoryId string) error {
	err := gateway.DeleteById(domain.CATEGORY, "id", categoryId)
	if err != nil {
		fmt.Println("error in usecase/DelCategoryById/gateway.DeleteById")
		return err
	}
	return nil
}

func UpsertCurriculum(id string, newData domain.CurriculumEntity, resourceType domain.ResourceType) error {
	err := gateway.UpsertBasePageById(
		resourceType,
		id,
		newData,
		func(c domain.CurriculumEntity) string { return c.Id },
	)
	if err != nil {
		fmt.Println("error in usecase/UpsertCurriculum/gateway.UpsertBasePageById")
		return err
	}
	return nil
}
func UpsertInfo(id string, newData domain.InfoEntity, resourceType domain.ResourceType) error {
	err := gateway.UpsertBasePageById(
		resourceType,
		id,
		newData,
		func(c domain.InfoEntity) string { return c.Id },
	)
	if err != nil {
		fmt.Println("error in usecase/UpsertCurriculum/gateway.UpsertBasePageById")
		return err
	}
	return nil
}
func UpsertAnswer(id string, newData domain.AnswerEntity, resourceType domain.ResourceType) error {
	err := gateway.UpsertBasePageById(
		resourceType,
		id,
		newData,
		func(c domain.AnswerEntity) string { return c.Id },
	)
	if err != nil {
		fmt.Println("error in usecase/UpsertCurriculum/gateway.UpsertBasePageById")
		return err
	}
	return nil
}

func UpsertSyncedFile(b domain.BlockEntity) error {
	err := gateway.UpsertSyncedDataById(domain.SYNCED, b.Id, b)
	if err != nil {
		fmt.Println("error in usecase/UpsertSyncedFile/gateway.UpsertSyncedDataById")
		return err
	}
	return nil
}

func UpsertChangesFile(changes_ []ChangeItem) error {
	_, err := gateway.UpsertFile(domain.CHANGES, "id", changes_)
	if err != nil {
		fmt.Println("error in usecase/fileUC.go: UpsertChangesFile/gateway/UpsertFile")
		return err
	}
	return nil
}

func InitCurriculumRelatedDir(curriculumId string) error {
	pageDatas, err := filemanager.ReadJson[[]domain.AtlPageEntity](constants.TMP_ALL_PAGE_PATH)
	if err != nil {
		fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.ReadJSONToMapArray")
		return err
	}
	var targetPages []domain.AtlPageEntity
	for _, item := range pageDatas {
		if curriculumId == item.CurriculumId {
			targetPages = append(targetPages, item)
		}
	}
	var changeItems []ChangeItem
	for _, page := range targetPages {
		err = filemanager.DelFile(fmt.Sprintf("%s/%s.dat", constants.PAGE_DATA_DIR, page.Id))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelFile")
			return err
		}
		err = filemanager.DelDir(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, page.Id))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelDir")
			return err
		}
		err = filemanager.DelFile(fmt.Sprintf("%s/%s.png", constants.OGP_DIR, page.Id))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelFile")
			return err
		}
		changeItem, err := NewChangeItem(page.Id, page.Title, "delete")
		if err != nil {
			fmt.Println("error in usecase/fileUC/fileUC.go: InitCurriculumRelatedDir/NewChangeItem")
			return err
		}
		changeItems = append(changeItems, *changeItem)
	}
	err = UpsertChangesFile(changeItems)
	if err != nil {
		fmt.Println("error in usecase/fileUC/fileUC.go: InitCurriculumRelatedDir/UpsertChangeFile")
		return err
	}
	fmt.Println("✅ completed: initialize curriculum related directory")
	return nil
}

func UpsertCategory(entities []domain.CategoryEntity) error {
	_, err := gateway.UpsertFile(domain.CATEGORY, "id", entities)
	if err != nil {
		fmt.Println("error in usecase/UpsertCategory/gateway.UpsertFile")
		return err
	}
	return nil
}
