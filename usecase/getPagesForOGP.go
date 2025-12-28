package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func GetPagesForOGP() ([]domain.PageEntity, error) {
	pagesP, err := GetPageFile()
	if err != nil {
		fmt.Println("error in presentation/HandleCreateGGP/GetPageFile")
		return nil, err
	}
	if pagesP == nil {
		return nil, fmt.Errorf("unexpected: pagesP is nil")
	}
	categoriesPointer, err := GetCategoryFile()
	if err != nil {
		fmt.Println("error in presentation/HandleCreateGGP/GetCategoryFile")
		return nil, err
	}
	if categoriesPointer == nil {
		return nil, fmt.Errorf("unexpected: categoriesPointer is nil")
	}
	pages := *pagesP
	var pageEntity []domain.PageEntity
	for _, p := range pages {
		ent, err := p.ToPageEntity()
		if err != nil {
			fmt.Println("error in usecase/getPagesForOGP.go/GetPagesForOGP/p.ToPageEntity")
			return nil, err
		}
		pageEntity = append(pageEntity, *ent)
	}
	categories := *categoriesPointer
	for _, c := range categories {
		entity, err := c.ToPageEntity()
		if err != nil {
			fmt.Println("error in usecase/GetPagesForOGPc.ToPageEntity/")
			return nil, err
		}
		pageEntity = append(pageEntity, *entity)
	}
	pageEntity = append(pageEntity, domain.CreatePage("部活情報", "emoji", "ℹ️", "infos"))
	pageEntity = append(pageEntity, domain.CreatePage("基礎班カリキュラム", "emoji", "🔰", "basic"))
	return pageEntity, nil
}
