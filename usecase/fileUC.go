package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
	"notion2atlas/utils"
	"strings"
)

func GetPageFile() (*[]domain.PageEntity, error) {
	data, err := gateway.GetFileData[[]domain.PageEntity](domain.PAGE)
	if err != nil {
		fmt.Println("error in usecase/GetPageFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func GetCurriculumFile() (*[]domain.CurriculumEntity, error) {
	data, err := gateway.GetFileData[[]domain.CurriculumEntity](domain.CURRICULUM)
	if err != nil {
		fmt.Println("error in usecase/GetCurriculumFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func GetCategoryFile() (*[]domain.CategoryEntity, error) {
	data, err := gateway.GetFileData[[]domain.CategoryEntity](domain.CATEGORY)
	if err != nil {
		fmt.Println("error in usecase/GetCurriculumFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func GetInfoFile() (*[]domain.InfoEntity, error) {
	data, err := gateway.GetFileData[[]domain.InfoEntity](domain.INFO)
	if err != nil {
		fmt.Println("error in usecase/GetInfoFile/gateway.GetFileData")
		return nil, err
	}
	return data, nil
}

func GetAnswerFile() (*[]domain.AnswerEntity, error) {
	data, err := gateway.GetFileData[[]domain.AnswerEntity](domain.ANSWER)
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
	noHyphenTargetId := strings.ReplaceAll(curriculumId, "-", "")
	pageDatas, err := filemanager.ReadJSONToMapArray(constants.PAGE_PATH)
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
		noHyphenId := strings.ReplaceAll(*id, "-", "")
		if noHyphenTargetId == noHyphenId {
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
		err = filemanager.DelFile(fmt.Sprintf("public/pageData/%s.json", *pageId))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelFile")
			return err
		}
		err = filemanager.DelDir("public/assets/" + *pageId)
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelDir")
			return err
		}
		err = filemanager.DelFile(fmt.Sprintf("public/ogp/%s.png", *pageId))
		if err != nil {
			fmt.Println("error in usecase/InitCurriculumRelatedDir/filemanager.DelFile")
			return err
		}
	}
	fmt.Println("✅ completed: initialize curriculum related directory")
	return nil
}

func UpsertCategory(entities []domain.CategoryEntity) error {
	err := gateway.UpsertFile(domain.CATEGORY, "id", entities)
	if err != nil {
		fmt.Println("error in usecase/UpsertCategory/gateway.UpsertFile")
		return err
	}
	return nil
}
