package presentation

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/usecase"
	"os"
)

func updateCategory() error {
	db_id := os.Getenv("NOTION_DB_ID_CATEGORY")
	publishedRecords, err := usecase.GetDBQuery(db_id)
	if err != nil {
		fmt.Println("error in presentation/updateCategory/usecase.GetDBQuery")
		return err
	}
	categories := []domain.CategoryEntity{}
	for _, record := range publishedRecords {
		category, err := record.ToCategoryEntity()
		if err != nil {
			fmt.Println("error in presentation/updateCategory/converter.Query2CCategoryEntity")
			return err
		}
		categories = append(categories, *category)
	}
	oldDataAddress, err := usecase.GetCategoryFile()
	if err != nil {
		fmt.Println("error in presentation/updateCategory/usecase.GetCurriculumFile")
		return err
	}
	if oldDataAddress == nil {
		return fmt.Errorf("oldDataAddress is nil")
	}
	oldData := *oldDataAddress
	nde, err := usecase.GetNDE(oldData, categories)
	if err != nil {
		fmt.Println("error in presentation/updateCategory/usecase.GetNDE")
		return err
	}
	fmt.Println("ℹ️ new", len(nde.New))
	fmt.Println("ℹ️ edited", len(nde.Edit))
	fmt.Println("ℹ️ delete", len(nde.Del))
	newCategory, err := domain.EntityIfArr2CategoryArr(nde.New)
	if err != nil {
		fmt.Println("error in presentation/updateCategory/domain.EntityIfArr2CategoryArr")
		return err
	}
	err = usecase.ProcessCategory(newCategory)
	if err != nil {
		fmt.Println("error in presentation/updateCategory/usecase.ProcessCategory")
		return err
	}
	editCategory, err := domain.EntityIfArr2CategoryArr(nde.Edit)
	if err != nil {
		fmt.Println("error in presentation/updateCategory/domain.EntityIfArr2CategoryArr")
		return err
	}
	err = usecase.ProcessCategory(editCategory)
	if err != nil {
		fmt.Println("error in presentation/updateCategory/usecase.ProcessCategory")
		return err
	}
	err = usecase.CreateStaticCategories()
	if err != nil {
		fmt.Println("error in presentation/updateCategory/usecase.CreateStaticCategories")
		return err
	}
	for _, id := range nde.Del {
		err = usecase.DelCategoryById(id)
		if err != nil {
			fmt.Println("error in presentation/updateCategory/usecase.DelCategoryById")
			return err
		}
		err = filemanager.DelFile(constants.OGP_DIR + id + ".png")
		if err != nil {
			fmt.Println("error in presentation/updateCategory/filemanager.DelFile")
			return err
		}
	}
	return nil
}
