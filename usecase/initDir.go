package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/filemanager"
)

func InitDir() error {
	require_dirs := [8]string{"public/assets", "public/pageData", constants.CURRICULUM_DIR, constants.PAGE_DIR, constants.CATEGORY_DIR, constants.INFO_DIR, constants.ANSWER_DIR, constants.SYNCED_DIR}
	for _, path := range require_dirs {
		_, err := filemanager.CreateDirIfNotExist(path)
		if err != nil {
			fmt.Println("error in filemanager/InitDir/CreateDirIfNotExist")
			return err
		}
	}
	require_files := [6]string{constants.CURRICULUM_PATH, constants.PAGE_PATH, constants.CATEGORY_PATH, constants.INFO_PATH, constants.ANSWER_PATH, constants.SYNCED_PATH}
	for _, path := range require_files {
		exists, err := filemanager.CreateFileIfNotExist(path)
		if err != nil {
			fmt.Println("error in usecase/InitDir/filemanager.CreateFileIfNotExist")
			return err
		}
		if !exists {
			err := filemanager.WriteJson([]any{}, path)
			if err != nil {
				fmt.Println("error in usecase/InitDir/filemanager.WriteJson")
				return err
			}
		}
	}
	return nil
}

func InitOGPDir() error {
	ogpDir := "public/ogp"
	_, err := filemanager.CreateDirIfNotExist(ogpDir)
	if err != nil {
		fmt.Println("error in usecase/InitOGPDir/filemanager.CreateDirIfNotExist")
		return err
	}
	err = filemanager.ClearDir(ogpDir)
	if err != nil {
		fmt.Println("error in usecase/InitOGPDir/filemanager.ClearDir")
		return err
	}
	return nil
}
