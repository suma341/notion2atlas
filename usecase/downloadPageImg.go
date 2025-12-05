package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"path"
)

type URLModel struct {
	IconUrl  string
	CoverUrl string
}

func DownloadPageImg(page domain.PageIf) error {
	iconType, iconUrl := page.GetIcon()
	coverType, coverUrl := page.GetCover()
	switch iconType {
	case "file", "custom_emoji":
		_, err := filemanager.DownloadFile(iconUrl, fmt.Sprintf("public/assets/%s", page.GetId()), "icon", ".png")
		if err != nil {
			fmt.Println("error in usecase/DownloadPageImg/filemanager.DownloadFile iconType")
			return err
		}
	}
	switch coverType {
	case "file":
		_, err := filemanager.DownloadFile(coverUrl, fmt.Sprintf("public/assets/%s", page.GetId()), "cover", ".png")
		if err != nil {
			fmt.Println("error in usecase/DownloadPageImg/filemanager.DownloadFile coverType")
			return err
		}
	}
	err := SaveOGPPicture(page)
	if err != nil {
		fmt.Println("error in usecase/DownloadPageImg/SaveOGPPicture")
		return err
	}
	return nil
}

func GetPathRewritedUrl(page domain.PageIf) URLModel {
	iconType, iconUrl := page.GetIcon()
	coverType, coverUrl := page.GetCover()
	var newIconUrl = ""
	var newCoverUrl = ""
	switch iconType {
	case "file", "custom_emoji":
		ext := path.Ext(iconUrl)
		if ext == "" {
			ext = ".png"
		}
		fileName := "icon" + ext
		newIconUrl = fmt.Sprintf("%s/assets/%s/%s", constants.DEPLOY_URL, page.GetId(), fileName)
	case "external", "emoji":
		newIconUrl = iconUrl
	}
	switch coverType {
	case "file":
		ext := path.Ext(coverUrl)
		if ext == "" {
			ext = ".png"
		}
		fileName := "cover" + ext
		newCoverUrl = fmt.Sprintf("%s/assets/%s/%s", constants.DEPLOY_URL, page.GetId(), fileName)
	case "external":
		newCoverUrl = coverUrl
	case "":
		newCoverUrl = ""
	}
	return URLModel{
		IconUrl:  newIconUrl,
		CoverUrl: newCoverUrl,
	}
}
