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
	// err := cleanTmpDir()
	// if err != nil {
	// 	fmt.Println("error in usecase/cleanup/cleanup.go:/CleanUp/cleanTmpDir")
	// 	return err
	// }
	return nil
}

func writeTestResult() error {
	page, err := filemanager.LoadAndDecodeJson[[]domain.AtlPageEntity](constants.PAGE_DAT_PATH)
	if err != nil {
		fmt.Println("❌ error in presentation/HandleUpdateData/filemanager.LoadAndDecodeJson")
		return err
	}
	err = filemanager.WriteJson(page, constants.TEST_RESULT_PATH)
	if err != nil {
		fmt.Println("❌ error in presentation/HandleUpdateData/filemanager.WriteJson")
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
