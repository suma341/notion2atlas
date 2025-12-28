package presentation

import (
	"fmt"
	"notion2atlas/usecase"
	postprocess "notion2atlas/usecase/PostProcess"
	"notion2atlas/usecase/cleanup"
	"notion2atlas/usecase/initapp"
)

func HandleUpdateData() error {
	var err error = nil
	err = initapp.InitApp()
	if err != nil {
		return err
	}
	fmt.Println("start process curriculums")
	curriculum_nde, err := updateCurriculum()
	if err != nil {
		return err
	}
	fmt.Println("start process Categories")
	err = updateCategory()
	if err != nil {
		return err
	}
	fmt.Println("start process infos")
	info_nde, err := updateInfo()
	if err != nil {
		return err
	}
	fmt.Println("start process answers")
	answer_nde, err := updateAnswer()
	if err != nil {
		return err
	}
	err = usecase.SaveStaticPageOGPPicture()
	if err != nil {
		return err
	}
	// var pageEntities = []domain.Entity{}
	pageEntities := append(
		curriculum_nde.New,
		curriculum_nde.Edit...,
	)
	pageEntities = append(pageEntities, info_nde.New...)
	pageEntities = append(pageEntities, info_nde.Edit...)
	pageEntities = append(pageEntities, answer_nde.New...)
	pageEntities = append(pageEntities, answer_nde.Edit...)
	err = postprocess.PostProcess()
	if err != nil {
		return err
	}
	err = cleanup.CleanUp()
	if err != nil {
		return err
	}
	return nil
}
