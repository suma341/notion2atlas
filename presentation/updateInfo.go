package presentation

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase"
	"os"
)

func updateInfo() error {
	db_id := os.Getenv("NOTION_DB_ID_INFO")
	publishedRecords, err := usecase.GetDBQuery(db_id)
	if err != nil {
		fmt.Println("error in presentation/UpdateInfo/usecase.GetDBQuery")
		return err
	}
	var infos []domain.InfoEntity
	for _, query := range publishedRecords {
		info, err := query.ToInfoEntity()
		if err != nil {
			fmt.Println("error in presentation/UpdateInfo/converter.Query2CurriculumEntity")
			return err
		}
		if info == nil {
			return fmt.Errorf("unexpected: curriculumEntity is nil")
		}
		infos = append(infos, *info)
	}
	oldDataAddress, err := usecase.GetInfoFile()
	if err != nil {
		fmt.Println("error in presentation/UpdateInfo/usecase.GetCurriculumFile")
		return err
	}
	if oldDataAddress == nil {
		return fmt.Errorf("oldDataAddress is nil")
	}
	oldData := *oldDataAddress
	err = usecase.ProcessNTData[domain.InfoEntity, domain.InfoEntity](oldData, infos, domain.INFO)
	if err != nil {
		fmt.Println("error in presentation/updateCurriculum/usecase.ProcessNTData")
		return err
	}
	return nil
}
