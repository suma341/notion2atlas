package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/usecase/fileUC"
	"notion2atlas/usecase/notionUC"
)

func saveNtBlockInPage(pageEntity domain.PageEntity, pageBuffer []domain.PageEntity, resourceType string) ([]domain.PageEntity, error) {
	filemanager.CreateDirIfNotExist(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, pageEntity.GetId()))
	err := filemanager.ClearDir(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, pageEntity.GetId()))
	if err != nil {
		fmt.Println("error in usecase/saveNtBlock.go: /saveNtBlockInPage/filemanager.ClearDir in curriculum/" + pageEntity.GetTitle())
		return pageBuffer, err
	}
	// urls, err := saveBasePage(pageEntity)
	// if err != nil {
	// 	fmt.Println("error in usecase/saveNtBlock.go: /saveNtBlockInPage/saveBasePage in curriculum/" + pageEntity.GetTitle())
	// 	return pageBuffer, err
	// }
	// urlRewritedEntity, err := pageEntity.ChangePageEntityUrl(urls.IconUrl, urls.CoverUrl)
	// if err != nil {
	// 	fmt.Println("error in usecase/saveNtBlock.go: /saveNtBlockInPage/basePage.ChangePageEntityUrl")
	// 	return pageBuffer, err
	// }
	// pageBuffer = append(pageBuffer, *urlRewritedEntity)
	blocks, err := notionUC.GetChildren(pageEntity.GetId())
	if err != nil {
		fmt.Println("error in usecase/saveNtBlock.go: /saveNtBlockInPage/GetChildren in curriculum/" + pageEntity.GetTitle())
		return nil, err
	}
	var blockBuffer = []domain.BlockEntity{}
	blockBuffer, pageBuffer, err = saveNtChildrenBlocks(pageEntity.GetId(), pageEntity.GetId(), blocks, blockBuffer, pageBuffer, resourceType)
	if err != nil {
		fmt.Println("error in usecase/saveNtBlock.go: /saveNtBlockInPage/saveNtChildrenBlocks in curriculum/" + pageEntity.GetTitle())
		return pageBuffer, err
	}
	err = FlushBlockBuffer(blockBuffer, pageEntity.GetId())
	if err != nil {
		fmt.Println("error in usecase/saveNtBlock.go: /saveNtBlockInPage/FlushBlockBuffer in curriculum/" + pageEntity.GetTitle())
		return pageBuffer, err
	}
	return pageBuffer, nil
}

func saveNtChildrenBlocks(currId string, pageId string, blocks []domain.NTBlockEntity, blockBuffer []domain.BlockEntity, pageBuffer []domain.PageEntity, resourceType string) ([]domain.BlockEntity, []domain.PageEntity, error) {
	var err error = nil
	for i, block := range blocks {
		blockBuffer, pageBuffer, err = saveNtBlock(block, currId, pageId, i, fmt.Sprintf("%d/%d", i+1, len(blocks)), blockBuffer, pageBuffer, resourceType)
		if err != nil {
			fmt.Println("error in usecase/saveNtChildrenBlocks/saveNtBlock")
			return blockBuffer, pageBuffer, err
		}
	}
	return blockBuffer, pageBuffer, nil
}

func saveNtBlock(block domain.NTBlockEntity, curriculumId string, pageId string, i int, p string, buffer []domain.BlockEntity, pageBuffer []domain.PageEntity, resourceType string) ([]domain.BlockEntity, []domain.PageEntity, error) {
	type_ := block.Type
	fmt.Println(p, type_)
	var err error = nil
	buffer, err = GetBlockEntities(block, buffer, curriculumId, pageId, i, resourceType)
	if err != nil {
		fmt.Println("error in usecase/saveNtBlock/GetBlockEntities in " + type_)
		return buffer, pageBuffer, err
	}
	if type_ == "child_page" {
		isNewPage_, err := isNewPage(block.Id)
		if err != nil {
			fmt.Println("error in usecase/saveNtBlock.go:/isNewPage/fileUC/GetPageFile")
			return buffer, pageBuffer, err
		}
		tmpPagePath := fmt.Sprintf("%s%s_new.json", constants.TMP_DIR, block.Id)
		if isNewPage_ {
			exist := filemanager.FileExists(tmpPagePath)
			if exist {
				pageEntity, err := filemanager.ReadJson[domain.PageEntity](tmpPagePath)
				if err != nil {
					fmt.Println("error in usecase/saveNtBlock.go:/isNewPage/filemanager.ReadJson")
					return buffer, pageBuffer, err
				}
				pageBuffer = append(pageBuffer, pageEntity)
			}
			filemanager.CreateDirIfNotExist(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, block.Id))
			children, err := notionUC.GetChildren(block.Id)
			if err != nil {
				fmt.Println("error in usecase/saveNtBlock/GetChildren in " + type_)
				return buffer, pageBuffer, err
			}
			var newBuffer = []domain.BlockEntity{}
			newBuffer, pageBuffer, err = saveNtChildrenBlocks(curriculumId, block.Id, children, newBuffer, pageBuffer, resourceType)
			if err != nil {
				fmt.Println("error in usecase/saveNtBlock/saveNtChildrenBlocks in " + type_)
				return buffer, pageBuffer, err
			}
			err = FlushBlockBuffer(newBuffer, block.Id)
			if err != nil {
				fmt.Println("error in usecase/saveNtBlock/FlushBlockBuffer in " + type_)
				return buffer, pageBuffer, err
			}
		}
		err = filemanager.DelFile(tmpPagePath)
		if err != nil {
			fmt.Println("error in usecase/saveNtBlock.go:/isNewPage/filemanager.DelFile")
			return buffer, pageBuffer, err
		}
	} else if (type_ == "synced_block" && block.SyncedBlock.SyncedFrom == nil) || type_ != "synced_block" {
		if block.HasChildren {
			children, err := notionUC.GetChildren(block.Id)
			if err != nil {
				fmt.Println("error in usecase/saveNtBlock/GetChildren in " + type_)
				return buffer, pageBuffer, err
			}
			buffer, pageBuffer, err = saveNtChildrenBlocks(curriculumId, pageId, children, buffer, pageBuffer, resourceType)
			if err != nil {
				fmt.Println("error in usecase/saveNtBlock/saveNtChildrenBlocks in " + type_)
				return buffer, pageBuffer, err
			}
		}
	}
	return buffer, pageBuffer, nil
}

func isNewPage(pageId string) (bool, error) {
	existPages, err := fileUC.GetPageFile()
	if err != nil {
		fmt.Println("error in usecase/saveNtBlock.go:/isNewPage/fileUC/GetPageFile")
		return false, err
	}
	for _, p := range *existPages {
		if pageId == p.Id {
			return false, nil
		}
	}
	return true, nil
}
