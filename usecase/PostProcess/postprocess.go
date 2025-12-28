package postprocess

import "fmt"

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
	err = encodeAndSaveSyncedDat()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/PostProcess/encodeAndSaveSyncedDat")
		return err
	}
	return nil
}
