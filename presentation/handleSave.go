package presentation

import (
	"fmt"
	"notion2atlas/usecase"
)

func HandleUpdateData() error {
	var err error = nil
	err = usecase.InitDir()
	if err != nil {
		panic(err)
	}
	fmt.Println("start process curriculums")
	err = updateCurriculum()
	if err != nil {
		return err
	}
	fmt.Println("start process Categories")
	err = updateCategory()
	if err != nil {
		return err
	}
	fmt.Println("start process infos")
	err = updateInfo()
	if err != nil {
		return err
	}
	fmt.Println("start process answers")
	err = updateAnswer()
	if err != nil {
		return err
	}
	err = usecase.SaveStaticPageOGPPicture()
	if err != nil {
		return err
	}
	return nil
}
