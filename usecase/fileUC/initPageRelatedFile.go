package fileUC

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/filemanager"
)

func InitPageRelatedFile(pageId string) error {
	err := filemanager.DelFile(fmt.Sprintf("%s/%s.dat", constants.PAGE_DATA_DIR, pageId))
	if err != nil {
		fmt.Println("error in usecase/InitPageRelatedFile/filemanager.DelFile")
		return err
	}
	err = filemanager.DelDir(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, pageId))
	if err != nil {
		fmt.Println("error in usecase/InitPageRelatedFile/filemanager.DelDir")
		return err
	}
	err = filemanager.DelFile(fmt.Sprintf("%s/%s.png", constants.OGP_DIR, pageId))
	if err != nil {
		fmt.Println("error in usecase/InitPageRelatedFile/filemanager.DelFile")
		return err
	}
	dirpath := fmt.Sprintf("%s/%s", constants.ASSETS_DIR, pageId)
	err = filemanager.ClearDir(dirpath)
	if err != nil {
		fmt.Println("error in usecase/InitPageRelatedFile/filemanager.ClearDir")
		return err
	}
	fmt.Println("✅ completed: initialize page related directory")
	return nil
}
