package postprocess

import (
	"errors"
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
	"os"
)

func processSyncedChild(atlEntity domain.AtlBlockEntity) (*[]domain.AtlBlockEntity, error) {
	syncedId := string(*atlEntity.Data.Synced)
	originalP, err := gateway.GetDatFileData[[]domain.BlockEntity](domain.SYNCED_DAT)
	if err != nil {
		fmt.Println("❌ error in postprocess/processSyncedChild/gateway.GetDatFileData")
		return nil, err
	}
	if originalP == nil {
		return nil, nil
	}
	originals := *originalP
	var original *domain.BlockEntity
	for _, ori := range originals {
		if syncedId == ori.Id {
			original = &ori
		}
	}
	if original == nil {
		fmt.Printf("original not found, synced_id:%s\n", syncedId)
		return nil, nil
	}
	pageDataPath := fmt.Sprintf("%s/%s.json", constants.PAGE_DATA_DIR, original.PageId)
	atlEntities, err := filemanager.ReadJson[[]domain.AtlBlockEntity](pageDataPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Println("❌ error in postprocess/processSyncedChild/filemanager.ReadJson:30")
			return nil, err
		}
		tmpDataPath := fmt.Sprintf("%s/%s.json", constants.TMP_DIR, original.PageId)
		entities, err := filemanager.ReadJson[[]domain.BlockEntity](tmpDataPath)
		if err != nil {
			fmt.Println("❌ error in postprocess/processSyncedChild/filemanager.ReadJson:37")
			return nil, err
		}
		if len(entities) == 0 {
			return nil, nil
		}
		atlEntities, err = processEntities(entities, original.Id, original.PageId)
		if err != nil {
			fmt.Println("❌ error in postprocess/processSyncedChild/processEntities:45")
			return nil, err
		}
		if len(atlEntities) == 0 {
			return nil, nil
		}
	}
	var oriChildren []domain.AtlBlockEntity
	for _, ent := range atlEntities {
		if syncedId == ent.ParentId {
			child := domain.AtlBlockEntity{
				Id:           ent.Id,
				Type:         ent.Type,
				ParentId:     atlEntity.Id,
				CurriculumId: atlEntity.CurriculumId,
				PageId:       atlEntity.PageId,
				Data:         ent.Data,
				Order:        ent.Order,
			}
			oriChildren = append(oriChildren, child)
			oriChildren = appendChild(oriChildren, atlEntities, child)
		}
	}
	return &oriChildren, nil
}

func appendChild(oriChildren []domain.AtlBlockEntity, entities []domain.AtlBlockEntity, parent domain.AtlBlockEntity) []domain.AtlBlockEntity {
	children := []domain.AtlBlockEntity{}
	for _, ent := range entities {
		if ent.ParentId == parent.Id {
			child := domain.AtlBlockEntity{
				Id:           ent.Id,
				Type:         ent.Type,
				ParentId:     parent.Id,
				CurriculumId: parent.CurriculumId,
				PageId:       parent.PageId,
				Data:         ent.Data,
				Order:        ent.Order,
			}
			children = append(children, child)
			children = appendChild(children, entities, child)
		}
	}
	return append(oriChildren, children...)
}

func processEntities(entities []domain.BlockEntity, originalId string, originalPageId string) ([]domain.AtlBlockEntity, error) {
	children := []domain.BlockEntity{}
	children = appendEntChild(children, entities, originalId)
	path := fmt.Sprintf("%s/%s.json", constants.TMP_DIR, originalPageId)
	pageEntities, err := filemanager.ReadJson[[]domain.PageEntity](path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Println("❌ error in postprocess/processEntities/filemanager.ReadJson:88")
			return nil, err
		}
		return nil, nil
	}
	atlBlocks, err := blockToAtlEntity(children, pageEntities)
	if err != nil {
		fmt.Println("❌ error in postprocess/processEntites/blockToAtlEntity:95")
		return nil, err
	}
	return atlBlocks, nil
}

func appendEntChild(children []domain.BlockEntity, allEntities []domain.BlockEntity, parentId string) []domain.BlockEntity {
	newChildren := []domain.BlockEntity{}
	for _, ent := range allEntities {
		if ent.ParentId == parentId {
			newChildren = append(newChildren, ent)
			newChildren = appendEntChild(newChildren, allEntities, ent.Id)
		}
	}
	return append(children, newChildren...)
}
