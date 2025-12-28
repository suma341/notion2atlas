package initapp

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
)

func initTest(test bool) error {
	if !test {
		err := filemanager.ClearDir(constants.TEST_DIR)
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.ClearDir")
			return err
		}
	} else {
		_, err := filemanager.CreateDirIfNotExist(constants.TEST_DIR)
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.CreateDirIfNotExist")
			return err
		}
		pageData, err := filemanager.LoadAndDecodeJson[[]domain.AtlPageEntity](constants.PAGE_DAT_PATH)
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.LoadAndDecodeJson")
			return err
		}
		err = filemanager.WriteJson(pageData, constants.TEST_PREV_PATH)
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.LoadAndDecodeJson")
			return err
		}
	}
	return nil
}
