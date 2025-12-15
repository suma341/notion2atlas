package usecase

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
)

func InsertNotionBlocks(bp domain.BasePage, pageBuffer []domain.PageEntity, resourceType string) ([]domain.PageEntity, error) {
	filemanager.CreateDirIfNotExist("public/assets/" + bp.GetId())
	err := filemanager.ClearDir("public/assets/" + bp.GetId())
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/filemanager.ClearDir in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	urls, pageEntity, err := insertBasePage(bp)
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/InsertBasePage in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	basePage := *pageEntity
	pageBuffer = append(pageBuffer, basePage.ChangePageEntityUrl(urls.IconUrl, urls.CoverUrl))
	blocks, err := GetChildren(bp.GetId())
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/GetChildren in curriculum/" + bp.GetTitle())
		return nil, err
	}
	var blockBuffer = []domain.BlockEntity{}
	blockBuffer, pageBuffer, err = insertChildren(bp.GetId(), bp.GetId(), blocks, blockBuffer, pageBuffer, resourceType)
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/insertChildren in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	err = FlushBlockBuffer(blockBuffer, bp.GetId())
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/FlushBlockBuffer in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	return pageBuffer, nil
}

func insertChildren(currId string, pageId string, blocks []domain.NTBlockRepository, blockBuffer []domain.BlockEntity, pageBuffer []domain.PageEntity, resourceType string) ([]domain.BlockEntity, []domain.PageEntity, error) {
	var err error = nil
	// var wg sync.WaitGroup
	// for i, block := range blocks {
	// 	block := block
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		blockBuffer, pageBuffer, err = insertBlock(block, currId, pageId, i, fmt.Sprintf("%d/%d", i+1, len(blocks)), blockBuffer, pageBuffer)
	// 		if err != nil {
	// 			fmt.Println("error in usecase/insertChildren/insertBlock " + block.Id)
	// 		}
	// 	}()
	// }
	for i, block := range blocks {
		blockBuffer, pageBuffer, err = insertBlock(block, currId, pageId, i, fmt.Sprintf("%d/%d", i+1, len(blocks)), blockBuffer, pageBuffer, resourceType)
		if err != nil {
			fmt.Println("error in usecase/insertChildren/insertBlock")
			return blockBuffer, pageBuffer, err
		}
	}

	// wg.Wait()
	return blockBuffer, pageBuffer, nil
}

func insertBlock(block domain.NTBlockRepository, curriculumId string, pageId string, i int, p string, buffer []domain.BlockEntity, pageBuffer []domain.PageEntity, resourceType string) ([]domain.BlockEntity, []domain.PageEntity, error) {
	type_ := block.Type
	fmt.Println(p, type_)
	var err error = nil
	buffer, pageBuffer, err = GetBlockEntities(block, buffer, curriculumId, pageId, i, pageBuffer, resourceType)
	if err != nil {
		fmt.Println("error in usecase/insertBlock/GetBlockEntities in " + type_)
		return buffer, pageBuffer, err
	}
	if type_ == "child_page" {
		filemanager.CreateDirIfNotExist("public/assets/" + block.Id)
		children, err := GetChildren(block.Id)
		if err != nil {
			fmt.Println("error in usecase/insertBlock/GetChildren in " + type_)
			return buffer, pageBuffer, err
		}
		var newBuffer = []domain.BlockEntity{}
		newBuffer, pageBuffer, err = insertChildren(curriculumId, block.Id, children, newBuffer, pageBuffer, resourceType)
		if err != nil {
			fmt.Println("error in usecase/insertBlock/insertChildren in " + type_)
			return buffer, pageBuffer, err
		}
		err = FlushBlockBuffer(newBuffer, block.Id)
		if err != nil {
			fmt.Println("error in usecase/insertBlock/FlushBlockBuffer in " + type_)
			return buffer, pageBuffer, err
		}
	} else if (type_ == "synced_block" && block.SyncedBlock.SyncedFrom == nil) || type_ != "synced_block" {
		if block.HasChildren {
			children, err := GetChildren(block.Id)
			if err != nil {
				fmt.Println("error in usecase/insertBlock/GetChildren in " + type_)
				return buffer, pageBuffer, err
			}
			buffer, pageBuffer, err = insertChildren(curriculumId, pageId, children, buffer, pageBuffer, resourceType)
			if err != nil {
				fmt.Println("error in usecase/insertBlock/insertChildren in " + type_)
				return buffer, pageBuffer, err
			}
		}
	}
	return buffer, pageBuffer, nil
}
