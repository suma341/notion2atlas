package presentation

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase"
)

func HandleUpdateData() error {
	var err error = nil
	err = usecase.InitDir()
	if err != nil {
		panic(err)
	}
	err = updateCurriculum()
	if err != nil {
		return err
	}
	err = updateCategory()
	if err != nil {
		return err
	}
	err = updateInfo()
	if err != nil {
		return err
	}
	err = updateAnswer()
	if err != nil {
		return err
	}
	staticPages := []domain.PageEntity{
		domain.CreatePage("部活情報", "emoji", "ℹ️", "infos"),
		domain.CreatePage("基礎班カリキュラム", "emoji", "🔰", "basic"),
	}
	for _, p := range staticPages {
		err := usecase.SaveOGPPicture(p)
		if err != nil {
			fmt.Println("error in presentation/HandleUpdateData/usecase.SaveOGPPicture")
			return err
		}
	}
	return nil
}
