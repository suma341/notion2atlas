package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
	"notion2atlas/utils"
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

func UpsertBasePage[T domain.BasePage](id string, newData T, resourceType domain.ResourceType) error {
	err := gateway.UpsertById(resourceType, id, newData)
	if err != nil {
		fmt.Println("error in usecase/UpsertCurriculum/gateway.UpsertById")
		return err
	}
	return nil
}

func UpsertSyncedFile(b domain.BlockEntity) error {
	err := gateway.UpsertById(domain.SYNCED, b.Id, b)
	if err != nil {
		fmt.Println("error in usecase/UpsertSyncedFile/gateway.UpsertById")
		return err
	}
	return nil
}

func InitCurriculumRelatedDir(curriculumId string) error {
	pageDatas, err := filemanager.ReadJSONToMapArray(constants.TMP_ALL_PAGE_PATH)
	if err != nil {
		fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.ReadJSONToMapArray")
		return err
	}
	var targetPages []map[string]any
	for _, item := range pageDatas {
		id, err := utils.SafelyRetrieve[string](item, "curriculumId")
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/utils.SafelyRetrieve")
			return err
		}
		if id == nil {
			return fmt.Errorf("unexpected: id is nil")
		}
		if curriculumId == *id {
			targetPages = append(targetPages, item)
		}
	}
	for _, page := range targetPages {
		pageId, err := utils.SafelyRetrieve[string](page, "id")
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/utils.SafelyRetrieve")
			return err
		}
		if pageId == nil {
			return fmt.Errorf("unexpected: pageId is nil")
		}
		err = filemanager.DelFile(fmt.Sprintf("%s/%s.dat", constants.PAGE_DATA_DIR, *pageId))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelFile")
			return err
		}
		err = filemanager.DelDir(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, *pageId))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelDir")
			return err
		}
		err = filemanager.DelFile(fmt.Sprintf("%s/%s.png", constants.OGP_DIR, *pageId))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelFile")
			return err
		}
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
