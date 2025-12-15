package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func ProcessCategory(categories []domain.CategoryEntity) error {
	processed := []domain.CategoryEntity{}
	for _, c := range categories {
		urls, _, err := insertBasePage(c)
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
		fmt.Println("error in presentation/updateCategory/usecase.UpsertCategory")
		return err
	}
	return nil
}
