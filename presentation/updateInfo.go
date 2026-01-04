package presentation

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase"
	"notion2atlas/usecase/fileUC"
	"notion2atlas/usecase/notionUC"
	"os"
)

func updateInfo() (*usecase.NDE, error) {
	db_id := os.Getenv("NOTION_DB_ID_INFO")
	publishedRecords, err := notionUC.GetDBQuery(db_id)
	if err != nil {
		fmt.Println("error in presentation/UpdateInfo/usecase.GetDBQuery")
		return nil, err
	}
	var infos []domain.InfoEntity
	for _, query := range publishedRecords {
		info, err := query.ToInfoEntity()
		if err != nil {
			fmt.Println("error in presentation/UpdateInfo/converter.Query2CurriculumEntity")
			return nil, err
		}
		if info == nil {
			return nil, fmt.Errorf("unexpected: curriculumEntity is nil")
		}
		infos = append(infos, *info)
	}
	oldDataAddress, err := fileUC.GetInfoFile()
	if err != nil {
		fmt.Println("error in presentation/UpdateInfo/usecase.GetCurriculumFile")
		return nil, err
	}
	if oldDataAddress == nil {
		return nil, fmt.Errorf("oldDataAddress is nil")
	}
	oldData := *oldDataAddress
	nde, err := usecase.ProcessNTData[domain.InfoEntity, domain.InfoEntity](oldData, infos, domain.INFO)
	if err != nil {
		fmt.Println("error in presentation/updateCurriculum/usecase.ProcessNTData")
		return nil, err
	}
	return nde, nil
}
