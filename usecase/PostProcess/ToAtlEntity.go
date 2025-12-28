package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
)

func rewriteToAtlEntity() error {
	pageEntities, err := filemanager.ReadJson[[]domain.PageEntity](constants.TMP_ALL_PAGE_PATH)
	if err != nil {
		fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/rewriteToAtlEntity/filemanager.ReadJson:11")
		return err
	}
	tmpPages, err := filemanager.ReadJson[[]domain.PageEntity](constants.TMP_PAGE_PATH)
	for _, page := range tmpPages {
		pageId := page.GetId()
		path := fmt.Sprintf("%s/%s.json", constants.TMP_DIR, pageId)
		blockEntities, err := filemanager.ReadJson[[]domain.BlockEntity](path)
		if err != nil {
			fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/rewriteToAtlEntity/filemanager.ReadJson:20")
			return err
		}
		atlBlocks, err := blockToAtlEntity(blockEntities, pageEntities)
		if err != nil {
			fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/rewriteToAtlEntity/blockToAtlEntity")
			return err
		}
		pageDataPath := fmt.Sprintf("%s/%s.dat", constants.PAGE_DATA_DIR, pageId)
		_, err = filemanager.CreateFileIfNotExist(pageDataPath)
		if err != nil {
			fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/rewriteToAtlEntity/filemanager.CreateFileIfNotExist")
			return err
		}
		err = filemanager.EncodeAndSave(atlBlocks, pageDataPath)
		// err = filemanager.WriteJson(atlBlocks, pageDataPath)
		if err != nil {
			fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/rewriteToAtlEntity/filemanager.WriteJson")
			return err
		}
		err = filemanager.DelFile(path)
		if err != nil {
			fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/rewriteToAtlEntity/filemanager.DelFile")
			return err
		}
	}
	fmt.Println("✅ complete process ntData to atlData")
	return nil
}

func blockToAtlEntity(blocks []domain.BlockEntity, pageEntities []domain.PageEntity) ([]domain.AtlBlockEntity, error) {
	atlBlocks := []domain.AtlBlockEntity{}
	for _, item := range blocks {
		var atlEntity domain.AtlBlockEntity
		switch item.Data.Type {
		case "paragraph", "todo", "header", "embed", "bookmark", "callout":
			atlEntity = processNormalParent(item, pageEntities)
		case "image":
			if item.Type == "image" {
				atlEntityPt, err := processImage(item, pageEntities)
				if err != nil {
					fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/blockToAtlEntity/processImage")
					return nil, err
				}
				if atlEntityPt == nil {
					continue
				}
				atlEntity = *atlEntityPt
			} else {
				atlEntity = processNormalParent(item, pageEntities)
			}
		case "table_row":
			atlEntity = processTableRowParent(item, pageEntities)
		case "code":
			atlEntity = processCodeParent(item, pageEntities)
		case "synced":
			if item.Data.Synced == nil || *item.Data.Synced == "" {
				continue
			}
			if *item.Data.Synced == "original" {
				atlData := domain.AtlBlockEntityData{
					Type:   item.Data.Type,
					Synced: item.Data.Synced,
				}
				atlBlocks = append(atlBlocks, item.ToAtlEntity(atlData))
			} else {
				atlEntity = item.ToAtlEntity(item.Data.ToAtlData(nil))
				childrenPt, err := processSyncedChild(atlEntity)
				if err != nil {
					fmt.Println("error in usecase/postprocess/ToAtlEntity.go:/blockToAtlEntity/processSyncedChild")
					return nil, err
				}
				if childrenPt != nil {
					atlChildren := *childrenPt
					atlBlocks = append(atlBlocks, atlChildren...)
				}
			}
		default:
			if item.Type == "table_of_contents" {
				data := processTOC(blocks)
				atlEntity = item.ToAtlEntity(data)
			} else {
				atlEntity = item.ToAtlEntity(item.Data.ToAtlData(nil))
			}
		}
		atlBlocks = append(atlBlocks, atlEntity)
	}
	return atlBlocks, nil
}

func getAllPagesInCurriculum(bps []domain.BasePage, pageEntities []domain.PageEntity) []domain.PageEntity {
	allPages := []domain.PageEntity{}
	for _, pe := range pageEntities {
		for _, bp := range bps {
			if pe.CurriculumId == bp.GetId() {
				allPages = append(allPages, pe)
			}
		}
	}
	return allPages
}

func processNormalParent(atl domain.BlockEntity, pageEntities []domain.PageEntity) domain.AtlBlockEntity {
	hasParentEntity := atl.Data.GetHasParentEntity()
	if hasParentEntity == nil {
		return atl.ToAtlEntity(atl.Data.ToAtlData(nil))
	}
	richTexts := hasParentEntity.GetParent()
	atlParents := processRichText(richTexts, pageEntities, atl.CurriculumId)
	atlData := atl.Data.ToAtlData(&atlParents)
	atlEntity := atl.ToAtlEntity(atlData)
	return atlEntity
}

func processCodeParent(atlItem domain.BlockEntity, pageEntities []domain.PageEntity) domain.AtlBlockEntity {
	code := atlItem.Data.Code
	cap := code.Caption
	atlCap := processRichText(cap, pageEntities, atlItem.CurriculumId)
	pa := code.Parent
	atlPa := processRichText(pa, pageEntities, atlItem.CurriculumId)
	atlCode := domain.AtlCodeEntity{
		Language: code.Language,
		Caption:  atlCap,
		Parent:   atlPa,
	}
	atlData := domain.AtlBlockEntityData{
		Type: atlItem.Type,
		Code: &atlCode,
	}
	atlBlock := atlItem.ToAtlEntity(atlData)
	return atlBlock
}

func processTableRowParent(atlItem domain.BlockEntity, pageEntities []domain.PageEntity) domain.AtlBlockEntity {
	atlParents := [][]domain.AtlRichTextEntity{}
	for _, blockArr := range *atlItem.Data.TableRow {
		atlParents = append(atlParents, processRichText(blockArr, pageEntities, atlItem.CurriculumId))
	}
	atlTableRow := atlItem.Data.TableRow.ToAtl(&atlParents)
	atlData := domain.AtlBlockEntityData{
		Type:     atlItem.Data.Type,
		TableRow: &atlTableRow,
	}
	atlBlock := atlItem.ToAtlEntity(atlData)
	return atlBlock
}
