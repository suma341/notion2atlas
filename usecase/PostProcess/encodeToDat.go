package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
)

func encodeAndSaveDats() error {
	err := encodeAndSaveSyncedDat()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/encodeAndSaveDats/encodeAndSaveSyncedDat")
		return err
	}
	err = encodeAndSaveCurriculumDat()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/encodeAndSaveDats/encodeAndSaveCurriculumDat")
		return err
	}
	err = encodeAndSaveCategoryDat()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/encodeAndSaveDats/encodeAndSaveCategoryDat")
		return err
	}
	err = encodeAndSaveInfoDat()
	if err != nil {
		fmt.Println("error in usecase/postprocess/postprocess.go:/encodeAndSaveDats/encodeAndSaveInfoDat")
		return err
	}
	return nil
}

func encodeAndSaveSyncedDat() error {
	return encodeAndSaveDatItem[[]domain.BlockEntity](
		domain.SYNCED_DAT,
		constants.SYNCED_DAT_PATH,
	)
}

func encodeAndSaveCurriculumDat() error {
	return encodeAndSaveDatItem[[]domain.CurriculumEntity](
		domain.CURRICULUM_DAT,
		constants.CURRICULUM_DAT_PATH,
	)
}

func encodeAndSaveCategoryDat() error {
	return encodeAndSaveDatItem[[]domain.CategoryEntity](
		domain.CATEGORY_DAT,
		constants.CATEGORY_DAT_PATH,
	)
}

func encodeAndSaveInfoDat() error {
	return encodeAndSaveDatItem[[]domain.InfoEntity](
		domain.INFO_DAT,
		constants.INFO_DAT_PATH,
	)
}

func encodeAndSaveDatItem[T any](datType domain.DatRType, savePath string) error {
	dat, err := gateway.GetDatFileData[T](datType)
	if err != nil {
		fmt.Println("❌ error in postprocess/encodeAndSaveDatItem/gateway.GetDatFileData")
		return err
	}

	if err := filemanager.EncodeAndSave(dat, savePath); err != nil {
		fmt.Println("❌ error in postprocess/encodeAndSaveDatItem/filemanager.EncodeAndSave")
		return err
	}

	return nil
}
