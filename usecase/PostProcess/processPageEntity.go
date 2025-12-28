package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
)

func processPageEntity() error {
	tmpPages, err := filemanager.ReadJson[[]domain.PageEntity](constants.TMP_PAGE_PATH)
	// filemanager.WriteJson(tmpPages, "notion_data/test.json")
	if err != nil {
		fmt.Println("error in usecase/postprocess/processPageEntity.go:/processPageEntity/filemanager.ReadJson")
		return err
	}
	curriculums, err := filemanager.ReadJson[[]domain.CurriculumEntity](constants.TMP_ALL_CURRICULUM_PATH)
	if err != nil {
		fmt.Println("error in usecase/postprocess/processPageEntity.go:/processPageEntity/filemanager.ReadJson")
		return err
	}
	infos, err := filemanager.ReadJson[[]domain.InfoEntity](constants.INFO_PATH)
	if err != nil {
		fmt.Println("error in usecase/postprocess/processPageEntity.go:/processPageEntity/filemanager.ReadJson")
		return err
	}
	answers, err := filemanager.ReadJson[[]domain.AnswerEntity](constants.ANSWER_PATH)
	if err != nil {
		fmt.Println("error in usecase/postprocess/processPageEntity.go:/processPageEntity/filemanager.ReadJson")
		return err
	}
	atlPageEntities := []domain.AtlPageEntity{}
	for _, page := range tmpPages {
		var filepath = fmt.Sprintf("%s/%s.json", constants.TMP_DIR, page.Id)
		blocks, err := filemanager.ReadJson[[]domain.BlockEntity](filepath)
		if err != nil {
			fmt.Println("error in usecase/postprocess/processPageEntity.go:/processPageEntity/filemanager.ReadJson")
			return err
		}
		firstText := getFirstText(blocks)
		imagePath := fmt.Sprintf("%s/ogp/%s.png", constants.DEPLOY_URL, page.Id)
		ogpData := domain.PageOgp{
			FirstText: firstText,
			ImagePath: imagePath,
		}
		var bpData BpData
		switch page.Type {
		case "curriculum":
			bpP := getBPData(curriculums, page.CurriculumId)
			if bpP == nil {
				continue
			}
			bpData = *bpP
		case "answer":
			bpP := getBPData(answers, page.CurriculumId)
			if bpP == nil {
				continue
			}
			bpData = *bpP
		case "info":
			bpP := getBPData(infos, page.CurriculumId)
			if bpP == nil {
				continue
			}
			bpData = *bpP
		}
		atlPageEntity := domain.NewAtlPageEntity(
			page.Id,
			page.CurriculumId,
			page.IconType,
			page.IconUrl,
			page.CoverUrl,
			page.CoverType,
			page.Order,
			page.ParentId,
			page.Title,
			page.Type,
			ogpData,
			bpData.Visibility,
			bpData.Tag,
			bpData.Category,
			page.LastEditedTime,
		)
		atlPageEntities = append(atlPageEntities, atlPageEntity)
	}
	newEntity, err := gateway.UpsertFile(domain.TMP_ALL_PAGE, "id", atlPageEntities)
	if err != nil {
		fmt.Println("error in usecase/postprocess/processPageEntity.go:/processPageEntity/gateway.UpsertFile")
		return err
	}
	err = filemanager.EncodeAndSave(newEntity, constants.PAGE_DAT_PATH)
	if err != nil {
		fmt.Println("error in usecase/postprocess/processPageEntity.go:/processPageEntity/filemanager.EncodeAndSave")
		return err
	}
	return nil
}

func getFirstText(blocks []domain.BlockEntity) string {
	firstText := ""
	for _, block := range blocks {
		switch block.Data.Type {
		case "paragraph":
			firstText = block.Data.Paragraph.GetConcatenatedText()
		case "todo":
			firstText = block.Data.Todo.GetConcatenatedText()
		case "header":
			firstText = block.Data.Header.GetCombinedPlainText()
		case "callout":
			firstText = block.Data.Callout.GetConcatenatedText()
		}
		if firstText != "" {
			break
		}
	}
	return firstText
}

type BpData struct {
	Visibility []string
	Tag        []string
	Category   []string
}

func getBPData[T domain.DBBasePage](bps []T, curriculumId string) *BpData {
	for _, bp := range bps {
		if bp.GetId() == curriculumId {
			data := BpData{
				Visibility: bp.GetVisilities(),
				Tag:        bp.GetTags(),
				Category:   bp.GetCategories(),
			}
			return &data
		}
	}
	return nil
}
