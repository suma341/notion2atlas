package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
)

func encodeAndSaveSyncedDat() error {
	syncedDat, err := gateway.GetDatFileData[[]domain.BlockEntity](domain.SYNCED_DAT)
	if err != nil {
		fmt.Println("❌ error in postprocess/encodeSyncedDat/gateway.GetDatFileData")
		return err
	}
	err = filemanager.EncodeAndSave(syncedDat, constants.SYNCED_DAT_PATH)
	if err != nil {
		fmt.Println("❌ error in postprocess/encodeSyncedDat/filemanager.EncodeAndSave")
		return err
	}
	return nil
}
