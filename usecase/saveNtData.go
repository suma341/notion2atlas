package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
)

func saveNtData(page domain.NtPageEntity, curriculumId string, existPages []domain.AtlPageEntity, resourceType domain.ResourceType) error {
	var err error = nil
	pageBuffer := []domain.PageEntity{}
	newPageEntity, err := getNewPageEntity(page, existPages, curriculumId)
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

func getNewPageEntity(page domain.NtPageEntity, existPages []domain.AtlPageEntity, curriculumId string) (*domain.PageEntity, error) {
	var target domain.AtlPageEntity
	for _, p := range existPages {
		if p.Id == page.Id {
			target = p
			break
		}
	}
	var newPageEntity domain.PageEntity
	if target.Id == "" || target.Order == 0 {
		isBp := page.Id == curriculumId
		if isBp {
			order := getBpOrder(page.Type, page.Id)
			fmt.Println("title: ", page.Title, " order: ", order)
			ent, err := domain.NewPageEntity(
				page.Id,
				curriculumId,
				page.IconType,
				page.IconUrl,
				page.CoverUrl,
				page.CoverType,
				order,
				"",
				page.Title,
				page.Type,
				page.LastEditedTime,
			)
			if err != nil {
				fmt.Println("error in usecase/saveNtData.go: getNewPageEntity/domain.NewPageEntity")
				return nil, err
			}
			newPageEntity = *ent
		} else {
			return nil, fmt.Errorf("unexpected: can not find page and is not base-page")
		}
	} else {
		newPageEntityP, err := domain.NewPageEntity(
			page.Id,
			curriculumId,
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
		if err != nil {
			fmt.Println("error in usecase/saveNtData.go: getNewPageEntity/domain.NewPageEntity")
			return nil, err
		}
		newPageEntity = *newPageEntityP
	}
	urls, err := saveBasePage(newPageEntity)
	if err != nil {
		fmt.Println("error in usecase/saveNtData.go: getNewPageEntity/saveBasePage in curriculum/" + newPageEntity.GetTitle())
		return nil, err
	}
	urlRewritedEntity, err := newPageEntity.ChangePageEntityUrl(urls.IconUrl, urls.CoverUrl)
	if err != nil {
		fmt.Println("error in usecase/saveNtData.go: getNewPageEntity/basePage.ChangePageEntityUrl")
		return nil, err
	}
	return urlRewritedEntity, nil
}

func getBpOrder(pageType string, pageId string) int {
	switch pageType {
	case "curriculum":
		curriculums, err := filemanager.ReadJson[[]domain.CurriculumEntity](constants.TMP_ALL_CURRICULUM_PATH)
		if err != nil {
			fmt.Println("❌ error in usecase/saveNtData.go: getBpOrder/filemanager.ReadJson")
			return 0
		}
		for _, c := range curriculums {
			if c.Id == pageId {
				return c.Order
			}
		}
	case "answer":
		ans, err := filemanager.ReadJson[[]domain.AnswerEntity](constants.TMP_ALL_ANSWER_PATH)
		if err != nil {
			fmt.Println("❌ error in usecase/saveNtData.go: getBpOrder/filemanager.ReadJson")
			return 0
		}
		for _, c := range ans {
			if c.Id == pageId {
				return c.Order
			}
		}
	case "info":
		inf, err := filemanager.ReadJson[[]domain.InfoEntity](constants.TMP_ALL_INFO_PATH)
		if err != nil {
			fmt.Println("❌ error in usecase/saveNtData.go: getBpOrder/filemanager.ReadJson")
			return 0
		}
		for _, c := range inf {
			if c.Id == pageId {
				return c.Order
			}
		}
	}
	return 0
}
