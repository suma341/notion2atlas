package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func ProcessCategory(categories []domain.CategoryEntity) error {
	processed := []domain.CategoryEntity{}
	for _, c := range categories {
		urls, _, err := saveBasePage(c)
		if err != nil {
			fmt.Println("error in presentation/updateCategory/usecase.InsertBasePage")
			return err
		}
		processed = append(processed, domain.CategoryEntity{
			Id:                c.Id,
			Title:             c.Title,
			Description:       c.Description,
			IconType:          c.IconType,
			IconUrl:           urls.IconUrl,
			CoverType:         c.CoverType,
			CoverUrl:          urls.CoverUrl,
			IsBasicCurriculum: c.IsBasicCurriculum,
			Order:             c.Order,
			LastEditedTime:    c.LastEditedTime,
			Update:            c.Update,
		})
		entity, err := c.ToPageEntity()
		if err != nil {
			fmt.Println("error in usecase/ProcessCategory/c.ToPageEntity")
			return err
		}
		err = SaveOGPPicture(entity)
		if err != nil {
			fmt.Println("error in usecase/ProcessCategory/SaveOGPPicture")
			return err
		}
	}
	err := UpsertCategory(processed)
	if err != nil {
		fmt.Println("error in usecase/updateCategory/usecase.UpsertCategory")
		return err
	}
	return nil
}

func CreateStaticCategories() error {
	var staticCategories []domain.CategoryEntity
	infoCategory := domain.CreateStaticCategory("info", "部活情報", "emoji", "ℹ️")
	answerCategory := domain.CreateStaticCategory("answer", "解答", "emoji", "✔️")
	staticCategories = append(staticCategories, []domain.CategoryEntity{infoCategory, answerCategory}...)
	err := UpsertCategory(staticCategories)
	if err != nil {
		fmt.Println("error in usecase/CreateStaticCategories/UpsertCategory")
		return err
	}
	return nil
}
