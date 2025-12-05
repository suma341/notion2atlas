package presentation

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase"
	"os"
)

func updateCurriculum() error {
	db_id := os.Getenv("NOTION_DB_ID_HORIZON")
	publishedRecords, err := usecase.GetDBQuery(db_id)
	if err != nil {
		fmt.Println("error in presentation/saveCurriculum/usecase.GetDBQuery")
		return err
	}
	var curriculums []domain.CurriculumEntity
	for _, query := range publishedRecords {
		curr, err := query.ToCurriculumEntity()
		if err != nil {
			fmt.Println("error in presentation/saveCurriculum/converter.Query2CurriculumEntity")
			return err
		}
		if curr == nil {
			return fmt.Errorf("unexpected: curriculumEntity is nil")
		}
		curriculums = append(curriculums, *curr)
	}
	oldDataAddress, err := usecase.GetCurriculumFile()
	if err != nil {
		fmt.Println("error in presentation/saveCurriculum/usecase.GetCurriculumFile")
		return err
	}
	if oldDataAddress == nil {
		return fmt.Errorf("oldDataAddress is nil")
	}
	oldData := *oldDataAddress
	err = usecase.ProcessNTData[domain.CurriculumEntity, domain.CurriculumEntity](oldData, curriculums, domain.CURRICULUM)
	if err != nil {
		fmt.Println("error in presentation/updateCurriculum/usecase.ProcessNTData")
		return err
	}

	fmt.Println("✅ completed: update curriculums")
	return nil
}
