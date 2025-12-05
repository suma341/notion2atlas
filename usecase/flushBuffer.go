package usecase

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
)

func FlushBlockBuffer(blocks []domain.BlockEntity, pageId string) error {
	path := fmt.Sprintf("public/pageData/%s.json", pageId)
	_, err := filemanager.CreateFileIfNotExist(path)
	if err != nil {
		fmt.Println("error in usecase/FlushBlockBuffer/filemanager.CreateFileIfNotExist")
		return err
	}
	err = filemanager.WriteJson(blocks, path)
	if err != nil {
		fmt.Println("error in usecase/FlushBlockBuffer/filemanager.WriteJson")
		return err
	}
	fmt.Println("✅ success flush block buffer in page:" + pageId)
	return nil
}

func FlushPageBuffer(pages []domain.PageEntity, curriculumId string) error {
	err := gateway.UpsertFile(domain.PAGE, "id", pages)
	if err != nil {
		fmt.Println("error in usecase/FlushPageBuffer/gateway.UpsertPage")
		return err
	}
	fmt.Println("✅ success flush page buffer in curriculum:" + curriculumId)
	return nil
}
