package presentation

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase"
	"notion2atlas/usecase/fileUC"
	"notion2atlas/usecase/notionUC"
	"os"
)

func updateAnswer() (*usecase.NDE, error) {
	db_id := os.Getenv("NOTION_DB_ID_ANSWER")
	publishedRecords, err := notionUC.GetDBQuery(db_id)
	if err != nil {
		fmt.Println("error in presentation/UpdateAnswer/usecase.GetDBQuery")
		return nil, err
	}
	var Answers []domain.AnswerEntity
	for _, query := range publishedRecords {
		Answer, err := query.ToAnswerEntity()
		if err != nil {
			fmt.Println("error in presentation/UpdateAnswer/converter.Query2CurriculumEntity")
			return nil, err
		}
		if Answer == nil {
			return nil, fmt.Errorf("unexpected: curriculumEntity is nil")
		}
		Answers = append(Answers, *Answer)
	}
	oldDataAddress, err := fileUC.GetAnswerFile()
	if err != nil {
		fmt.Println("error in presentation/UpdateAnswer/usecase.GetCurriculumFile")
		return nil, err
	}
	if oldDataAddress == nil {
		return nil, fmt.Errorf("oldDataAddress is nil")
	}
	oldData := *oldDataAddress
	nde, err := usecase.ProcessNTData[domain.AnswerEntity, domain.AnswerEntity](oldData, Answers, domain.ANSWER)
	if err != nil {
		fmt.Println("error in presentation/UpdateAnswer/usecase.ProcessNTData")
		return nil, err
	}
	return nde, nil
}
