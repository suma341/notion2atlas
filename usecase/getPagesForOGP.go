package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func GetPagesForOGP() ([]domain.PageEntity, error) {
	pagesPointer, err := GetPageFile()
	if err != nil {
		fmt.Println("error in presentation/HandleCreateGGP/GetPageFile")
		return nil, err
	}
	if pagesPointer == nil {
		return nil, fmt.Errorf("unexpected: pagesPointer is nil")
	}
	categoriesPointer, err := GetCategoryFile()
	if err != nil {
		fmt.Println("error in presentation/HandleCreateGGP/GetCategoryFile")
		return nil, err
	}
	if categoriesPointer == nil {
		return nil, fmt.Errorf("unexpected: categoriesPointer is nil")
	}
	pages := *pagesPointer
	categories := *categoriesPointer
	for _, c := range categories {
		pages = append(pages, c.ToPageEntity())
	}
	pages = append(pages, domain.CreatePage("部活情報", "emoji", "ℹ️", "infos"))
	pages = append(pages, domain.CreatePage("基礎班カリキュラム", "emoji", "🔰", "basic"))
	return pages, nil
}
