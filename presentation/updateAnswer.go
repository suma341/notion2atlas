package presentation

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase"
	"os"
)

func updateAnswer() error {
	db_id := os.Getenv("NOTION_DB_ID_ANSWER")
	publishedRecords, err := usecase.GetDBQuery(db_id)
	if err != nil {
		fmt.Println("error in presentation/UpdateAnswer/usecase.GetDBQuery")
		return err
	}
	var Answers []domain.AnswerEntity
	for _, query := range publishedRecords {
		Answer, err := query.ToAnswerEntity()
		if err != nil {
			fmt.Println("error in presentation/UpdateAnswer/converter.Query2CurriculumEntity")
			return err
		}
		if Answer == nil {
			return fmt.Errorf("unexpected: curriculumEntity is nil")
		}
		Answers = append(Answers, *Answer)
	}
	oldDataAddress, err := usecase.GetAnswerFile()
	if err != nil {
		fmt.Println("error in presentation/UpdateAnswer/usecase.GetCurriculumFile")
		return err
	}
	if oldDataAddress == nil {
		return fmt.Errorf("oldDataAddress is nil")
	}
	oldData := *oldDataAddress
	err = usecase.ProcessNTData[domain.AnswerEntity, domain.AnswerEntity](oldData, Answers, domain.ANSWER)
	if err != nil {
		fmt.Println("error in presentation/UpdateAnswer/usecase.ProcessNTData")
		return err
	}
	return nil
}
