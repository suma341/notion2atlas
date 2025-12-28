package initapp

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
)

func initDir() error {
	require_dirs := [9]string{constants.ASSETS_DIR, constants.PAGE_DATA_DIR, constants.CURRICULUM_DIR, constants.PAGE_DIR, constants.CATEGORY_DIR, constants.INFO_DIR, constants.ANSWER_DIR, constants.TMP_DIR, constants.SYNCED_DIR}
	for _, path := range require_dirs {
		_, err := filemanager.CreateDirIfNotExist(path)
		if err != nil {
			fmt.Println("error in filemanager/InitDir/CreateDirIfNotExist")
			return err
		}
	}
	require_files := [2]string{constants.ANSWER_PATH, constants.TMP_PAGE_PATH}
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
	err := loadDat()
	if err != nil {
		fmt.Println("error in usescase/initprocess/initDir.go:/InitDir/loadDat")
		return err
	}
	return nil
}

func loadDat() error {
	err := loadAndWrite[[]domain.AtlPageEntity](constants.PAGE_DAT_PATH, constants.TMP_ALL_PAGE_PATH)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initDir.go:/loadDat/loadAndWrite")
		return err
	}
	err = loadAndWrite[[]domain.BlockEntity](constants.SYNCED_DAT_PATH, constants.TMP_ALL_SYNCED_PATH)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initDir.go:/loadDat/loadAndWrite")
		return err
	}
	err = loadAndWrite[[]domain.CurriculumEntity](constants.CURRICULUM_DAT_PATH, constants.TMP_ALL_CURRICULUM_PATH)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initDir.go:/loadDat/loadAndWrite")
		return err
	}
	err = loadAndWrite[[]domain.CategoryEntity](constants.CATEGORY_DAT_PATH, constants.TMP_ALL_CATEGORY_PATH)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initDir.go:/loadDat/loadAndWrite")
		return err
	}
	err = loadAndWrite[[]domain.InfoEntity](constants.INFO_DAT_PATH, constants.TMP_ALL_INFO_PATH)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initDir.go:/loadDat/loadAndWrite")
		return err
	}
	return nil
}

func loadAndWrite[T any](datPath string, tmpAllPath string) error {
	data, err := filemanager.LoadAndDecodeJson[T](datPath)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initDir.go:/loadDat/filemanager.LoadAndDecodeJson")
		return err
	}
	err = filemanager.WriteJson(data, tmpAllPath)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initDir.go:/loadDat/filemanager.WriteJson")
		return err
	}
	return err
}
