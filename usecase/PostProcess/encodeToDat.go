package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
)

func encodeAndSaveSyncedDat() error {
	return encodeAndSaveDat[[]domain.BlockEntity](
		domain.SYNCED_DAT,
		constants.SYNCED_DAT_PATH,
	)
}

func encodeAndSaveCurriculumDat() error {
	return encodeAndSaveDat[[]domain.CurriculumEntity](
		domain.CURRICULUM_DAT,
		constants.CURRICULUM_DAT_PATH,
	)
}

func encodeAndSaveDat[T any](datType domain.DatRType, savePath string) error {
	dat, err := gateway.GetDatFileData[T](datType)
	if err != nil {
		fmt.Println("❌ error in postprocess/encodeAndSaveDat/gateway.GetDatFileData")
		return err
	}

	if err := filemanager.EncodeAndSave(dat, savePath); err != nil {
		fmt.Println("❌ error in postprocess/encodeAndSaveDat/filemanager.EncodeAndSave")
		return err
	}

	return nil
}
