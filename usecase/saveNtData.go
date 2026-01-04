package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func saveNtData(page domain.NtPageEntity, curriculumId string, existPages []domain.AtlPageEntity, resourceType domain.ResourceType) error {
	var err error = nil
	pageBuffer := []domain.PageEntity{}
	newPageEntity, err := getNewPageEntity(page, existPages)
	if err != nil {
		fmt.Println("error in usecase/saveNtData.go: saveNtData/getNewPageEntity")
		return err
	}
	pageBuffer = append(pageBuffer, *newPageEntity)
	pageBuffer, err = saveNtBlockInPage(*newPageEntity, pageBuffer, resourceType.GetStr())
	if err != nil {
		fmt.Println("error in usecase/saveNtData/InsertCurriculumBlocks New")
		return err
	}
	err = FlushPageBuffer(pageBuffer, curriculumId)
	if err != nil {
		fmt.Println("error in usecase/saveNtData/FlushPageBuffer")
		return err
	}
	fmt.Println("✅ complete read : " + page.GetTitle())
	return nil
}

func getNewPageEntity(page domain.NtPageEntity, existPages []domain.AtlPageEntity) (*domain.PageEntity, error) {
	var target domain.AtlPageEntity
	for _, p := range existPages {
		if p.Id == page.Id {
			target = p
			break
		}
	}
	newPageEntity, err := domain.NewPageEntity(
		page.Id,
		target.CurriculumId,
		page.IconType,
		page.IconUrl,
		page.CoverUrl,
		page.CoverType,
		target.Order,
		target.ParentId,
		page.Title,
		page.Type,
		page.LastEditedTime,
	)
	urls, err := saveBasePage(*newPageEntity)
	if err != nil {
		fmt.Println("error in usecase/saveNtData.go: /getNewPageEntity/saveBasePage in curriculum/" + newPageEntity.GetTitle())
		return nil, err
	}
	urlRewritedEntity, err := newPageEntity.ChangePageEntityUrl(urls.IconUrl, urls.CoverUrl)
	if err != nil {
		fmt.Println("error in usecase/saveNtData.go: /getNewPageEntity/basePage.ChangePageEntityUrl")
		return nil, err
	}
	return urlRewritedEntity, nil
}
