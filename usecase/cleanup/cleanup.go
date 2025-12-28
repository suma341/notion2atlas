package cleanup

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"os"
)

func CleanUp() error {
	test := os.Getenv("TEST")
	is_test := test == "true"
	if is_test {
		err := writeTestResult()
		if err != nil {
			fmt.Println("error in usecase/cleanup/cleanup.go:/CleanUp/writeTestResult")
			return err
		}
	}
	err := cleanTmpDir()
	if err != nil {
		fmt.Println("error in usecase/cleanup/cleanup.go:/CleanUp/cleanTmpDir")
		return err
	}
	return nil
}

func writeTestResult() error {
	err := loadAndWrite[[]domain.AtlPageEntity](constants.PAGE_DAT_PATH, constants.TEST_RESULT_PAGE_PATH)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/writeTestResult/loadAndwrite")
		return err
	}
	err = loadAndWrite[[]domain.CurriculumEntity](constants.CURRICULUM_DAT_PATH, constants.TEST_RESULT_CURRICULUM_PATH)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/writeTestResult/loadAndwrite")
		return err
	}
	err = loadAndWrite[[]domain.CategoryEntity](constants.CATEGORY_DAT_PATH, constants.TEST_RESULT_CATEGORY_PATH)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/writeTestResult/loadAndwrite")
		return err
	}
	err = loadAndWrite[[]domain.InfoEntity](constants.INFO_DAT_PATH, constants.TEST_RESULT_INFO_PATH)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/writeTestResult/loadAndwrite")
		return err
	}
	err = loadAndWrite[[]domain.AnswerEntity](constants.ANSWER_DAT_PATH, constants.TEST_RESULT_ANSWER_PATH)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/writeTestResult/loadAndwrite")
		return err
	}
	err = loadAndWrite[[]domain.BlockEntity](constants.SYNCED_DAT_PATH, constants.TEST_RESULT_SYNCED_PATH)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/writeTestResult/loadAndwrite")
		return err
	}
	return nil
}

func loadAndWrite[T any](datPath string, resultPath string) error {
	data, err := filemanager.LoadAndDecodeJson[T](datPath)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/loadAndWrite/filemanager.LoadAndDecodeJson")
		return err
	}
	err = filemanager.WriteJson(data, resultPath)
	if err != nil {
		fmt.Println("error in usecase/cleanup.go:/loadAndWrite/filemanager.WriteJson")
		return err
	}
	return nil
}

func cleanTmpDir() error {
	err := filemanager.ClearDir(constants.TMP_DIR)
	if err != nil {
		fmt.Println("❌ error in postprocess/addOgpDataToPage/filemanager.ClearDir")
		return err
	}
	return nil
}
