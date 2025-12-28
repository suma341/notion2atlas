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
		err := createTestDir()
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/createTestDir")
			return err
		}
		err = createTestFile[[]domain.AtlPageEntity](constants.PAGE_DAT_PATH, constants.TEST_PREV_PAGE_PATH)
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.LoadAndDecodeJson")
			return err
		}
		err = createTestFile[[]domain.CurriculumEntity](constants.CURRICULUM_DAT_PATH, constants.TEST_PREV_CURRICULUM_PATH)
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.LoadAndDecodeJson")
			return err
		}
	}
	return nil
}

func createTestFile[T any](datPath string, testPrevPath string) error {
	data, err := filemanager.LoadAndDecodeJson[T](datPath)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.LoadAndDecodeJson")
		return err
	}
	_, err = filemanager.CreateFileIfNotExist(testPrevPath)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.CreateFileIfNotExist")
		return err
	}
	err = filemanager.WriteJson(data, testPrevPath)
	if err != nil {
		fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.LoadAndDecodeJson")
		return err
	}
	return nil
}

func createTestDir() error {
	testDir := [3]string{constants.TEST_DIR, constants.TEST_PREV_DIR, constants.TEST_RESULT_DIR}
	for _, dir := range testDir {
		_, err := filemanager.CreateDirIfNotExist(dir)
		if err != nil {
			fmt.Println("error in usecase/initprocess/initTest.go:/InitTest/filemanager.CreateDirIfNotExist")
			return err
		}
	}
	return nil
}
