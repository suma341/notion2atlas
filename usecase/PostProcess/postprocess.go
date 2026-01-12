package postprocess

import (
	"fmt"
)

func PostProcess() error {
	err := processPageEntity()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/PostProcess/addOgpDataToPage")
		return err
	}
	err = rewriteToAtlEntity()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/PostProcess/rewriteToAtlEntity")
		return err
	}
	err = encodeAndSaveDats()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/PostProcess/encodeAndSaveDats")
		return err
	}
	changeContents, err := getChanges()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/PostProcess/getChanges")
		return err
	}
	err = updateVersion(changeContents.isChanged())
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go: PostProcess/createChangesMessage")
		return err
	}
	message, err := createChangesMessage(*changeContents)
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go: PostProcess/createChangesMessage")
		return err
	}
	err = SendDiscordMessage(*message)
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go: PostProcess/sendDiscordMessage")
		return err
	}
	return nil
}
