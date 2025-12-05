package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func SaveNotionData[T domain.BasePage](item domain.Entity, resourceType domain.ResourceType) error {
	var err error = nil
	curr := item.(T)
	pageBuffer := []domain.PageEntity{}
	pageBuffer, err = InsertNotionBlocks(curr, pageBuffer)
	if err != nil {
		fmt.Println("error in usecase/SaveNotionData/InsertCurriculumBlocks New")
		return err
	}
	err = FlushPageBuffer(pageBuffer, curr.GetId())
	if err != nil {
		fmt.Println("error in usecase/SaveNotionData/FlushPageBuffer")
		return err
	}
	err = UpsertBasePage(curr.GetId(), curr, resourceType)
	if err != nil {
		fmt.Println("error in usecase/SaveNotionData/UpsertCurriculum")
		return err
	}
	fmt.Println("✅ complete read : " + curr.GetTitle())
	return nil
}
