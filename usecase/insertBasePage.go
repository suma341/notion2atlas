package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func insertBasePage(bp domain.BasePage) (*URLModel, *domain.PageEntity, error) {
	pageEntity := bp.ToPageEntity()
	err := DownloadPageImg(pageEntity)
	if err != nil {
		fmt.Println("error in usecase/InsertBasePage/DownloadPageImg")
		return nil, nil, err
	}
	urls := GetPathRewritedUrl(pageEntity)
	return &urls, &pageEntity, nil
}
