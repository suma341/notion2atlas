package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func saveBasePage(pageEntity domain.PageEntity) (*URLModel, error) {
	err := DownloadPageImg(pageEntity)
	if err != nil {
		fmt.Println("error in usecase/saveBasePage/DownloadPageImg")
		return nil, err
	}
	urls := GetPathRewritedUrl(pageEntity)
	return &urls, nil
}
