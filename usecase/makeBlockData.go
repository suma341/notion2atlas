package usecase

import (
	"errors"
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/utils"
	"strings"
)

func makeParagraphData(para domain.ParagraphProperty, type_ string) (*domain.BlockEntityData, error) {
	richTextModels, err := ProcessRichText(para.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeParagraphData/ProcessRichText")
		return nil, err
	}
	paragraph := domain.ParagraphEntity{
		Color:  para.Color,
		Parent: richTextModels,
	}
	data := domain.BlockEntityData{
		Type:      paragraph.GetType(),
		Paragraph: &paragraph,
	}
	return &data, nil
}

func makeToDoData(todo domain.ToDoProperty, type_ string) (*domain.BlockEntityData, error) {
	richTextModels, err := ProcessRichText(todo.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeParagraphData/ProcessRichText")
		return nil, err
	}
	todoEntity := domain.TodoEntity{
		Color:   todo.Color,
		Parent:  richTextModels,
		Checked: todo.Checked,
	}
	data := domain.BlockEntityData{
		Type: todoEntity.GetType(),
		Todo: &todoEntity,
	}
	return &data, nil
}

func makeHeaderData(header domain.HeaderProperty, type_ string) (*domain.BlockEntityData, error) {
	richTextModels, err := ProcessRichText(header.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeHeaderData/ProcessRichText")
		return nil, err
	}
	headerEntity := domain.HeaderEntity{
		Color:        header.Color,
		Parent:       richTextModels,
		IsToggleable: header.IsToggleable,
	}
	data := domain.BlockEntityData{
		Type:   headerEntity.GetType(),
		Header: &headerEntity,
	}
	return &data, nil
}

func makeImageData(img domain.ImageProperty, blockId string, pageId string, type_ string) (*domain.BlockEntityData, error) {
	richTextModels, err := ProcessRichText(img.Caption, type_)
	if err != nil {
		fmt.Println("error in usecase/makeImageData/ProcessRichText")
		return nil, err
	}
	path := ""
	if img.File != nil {
		fileName, err := filemanager.DownloadFile(img.File.Url, fmt.Sprintf("%s/%s", constants.ASSETS_DIR, pageId), blockId, ".png")
		if err != nil {
			fmt.Println("error in usecase/makeImageData/filemanager.DownloadFile")
			return nil, err
		}
		path = utils.GetDownloadPath(pageId, fileName)
	} else if img.External != nil {
		path = img.External.Url
	} else {
		fmt.Println("ℹ️ unexpected type: " + img.Type)
		filemanager.WriteJson(map[string]any{"type": img.Type}, "notion_data/test.json")
	}
	imageEntity := domain.ImageEntity{
		Parent: richTextModels,
		Url:    path,
	}
	data := domain.BlockEntityData{
		Type:  imageEntity.GetType(),
		Image: &imageEntity,
	}

	return &data, nil
}

func makeEmbedData(embed domain.EmbedProperty, type_ string) (*domain.BlockEntityData, error) {
	var data domain.BlockEntityData
	parent, err := ProcessRichText(embed.Caption, type_)
	if err != nil {
		fmt.Println("error in usecase/makeEmbedData/ProcessRichText")
		return nil, err
	}
	switch type_ {
	case "embed":
		canEmbed := utils.CanEmbed(embed.Url)
		embedEntity := domain.EmbedEntity{
			CanEmbed: canEmbed,
			Parent:   parent,
			Url:      embed.Url,
		}
		data = domain.BlockEntityData{
			Type:  embedEntity.GetType(),
			Embed: &embedEntity,
		}
	case "bookmark":
		ogpData, err := utils.GetOGP(embed.Url)
		if err != nil {
			fmt.Println("error in usecase/makeEmbedData/utils.GetOGP")
			return nil, err
		}
		bookmarkEntity := domain.BookmarkEntity{
			Parent: parent,
			Url:    embed.Url,
			Ogp:    *ogpData,
		}
		data = domain.BlockEntityData{
			Type:     bookmarkEntity.GetType(),
			Bookmark: &bookmarkEntity,
		}
	}
	return &data, nil
}

func makeTableData(table domain.TableProperty) (*domain.BlockEntityData, error) {
	tableEntity := domain.TableEntity(table)
	data := domain.BlockEntityData{
		Type:  "table",
		Table: &tableEntity,
	}
	return &data, nil
}

func makeTableRowData(table_row domain.TableRowProperty, type_ string) (*domain.BlockEntityData, error) {
	var cells [][]domain.RichTextEntity
	for _, cell := range table_row.Cells {
		richTextModel, err := ProcessRichText(cell, type_)
		if err != nil {
			fmt.Println("error in usecase/makeTableRowData/ProcessRichText")
			return nil, err
		}
		cells = append(cells, richTextModel)
	}
	tableRowEntity := domain.TableRowEntity(cells)
	data := domain.BlockEntityData{
		Type:     tableRowEntity.GetType(),
		TableRow: &tableRowEntity,
	}
	return &data, nil
}

func makeChildPageData(pageId string, type_ string) (*domain.BlockEntityData, *domain.NtPageEntity, error) {
	pageDataAddress, err := GetPageItem(pageId, type_)
	if err != nil {
		if errors.Is(err, domain.ErrNotionErrorResponse) {
			return nil, nil, domain.ErrNotionErrorResponse
		}
		fmt.Println("error in usecase/makeChildPageData/GetPageItem")
		return nil, nil, err
	}
	if pageDataAddress == nil {
		return nil, nil, fmt.Errorf("pageDataAddress is nil")
	}
	err = DownloadPageImg(pageDataAddress)
	if err != nil {
		fmt.Println("error in usecase/makeChildPageData/DownloadPageImg")
		return nil, nil, err
	}
	urls := GetPathRewritedUrl(pageDataAddress)
	childPageEntity := domain.ChildPageEntity{
		Parent:   pageDataAddress.Title,
		IconType: pageDataAddress.IconType,
		IconUrl:  urls.IconUrl,
		CoverUrl: urls.CoverUrl,
	}
	data := domain.BlockEntityData{
		Type:      childPageEntity.GetType(),
		ChildPage: &childPageEntity,
	}
	pageRepo := domain.NewNtPageEntity(
		pageDataAddress.Id,
		urls.IconUrl,
		pageDataAddress.IconType,
		urls.CoverUrl,
		pageDataAddress.CoverType,
		pageDataAddress.Title,
		pageDataAddress.Type,
		pageDataAddress.LastEditedTime,
		pageDataAddress.InTrash,
	)
	return &data, &pageRepo, nil
}

func makeLinkToPageData(link_to_page domain.LinkToPageProperty, type_ string) (*domain.BlockEntityData, error) {
	pageId := link_to_page.PageId
	nohyphenId := strings.ReplaceAll(pageId, "-", "")
	link := "/posts/curriculums/" + nohyphenId
	pageData, err := GetPageItem(pageId, type_)
	if err != nil {
		if errors.Is(err, domain.ErrNotionErrorResponse) {
			return nil, domain.ErrNotionErrorResponse
		}
		fmt.Println("error in usecase/makeLinkToPageData/GetPageItem")
		return nil, err
	}
	urls := GetPathRewritedUrl(pageData)
	linktopageEntity := domain.LinkToPageEntity{
		Link:     link,
		IconUrl:  urls.IconUrl,
		IconType: pageData.IconType,
		Title:    pageData.Title,
	}
	data := domain.BlockEntityData{
		Type:       linktopageEntity.GetType(),
		LinkToPage: &linktopageEntity,
	}
	return &data, nil
}

func makeCodeData(code domain.CodeProperty, type_ string) (*domain.BlockEntityData, error) {
	codeContent, err := ProcessRichText(code.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeCodeData/ProcessRichText(code.RichText)")
		return nil, err
	}
	caption, err := ProcessRichText(code.Caption, type_)
	if err != nil {
		fmt.Println("error in usecase/makeCodeData/ProcessRichText(code.Caption)")
		return nil, err
	}
	codeEntity := domain.CodeEntity{
		Language: code.Language,
		Caption:  caption,
		Parent:   codeContent,
	}
	data := domain.BlockEntityData{
		Type: codeEntity.GetType(),
		Code: &codeEntity,
	}
	return &data, nil
}

func makeCalloutData(callout domain.CalloutProperty, type_ string) (*domain.BlockEntityData, error) {
	richText, err := ProcessRichText(callout.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeCalloutData/ProcessRichText")
		return nil, err
	}
	calloutEntity := domain.CalloutEntity{
		Icon:   callout.Icon,
		Color:  callout.Color,
		Parent: richText,
	}
	data := domain.BlockEntityData{
		Type:    calloutEntity.GetType(),
		Callout: &calloutEntity,
	}
	return &data, nil
}

func makeSyncedBlockData(syncedBlock domain.SyncedProperty) *domain.BlockEntityData {
	syncedFrom := syncedBlock.SyncedFrom
	if syncedFrom == nil {
		syncedEntity := domain.SyncedEntity("original")
		data := domain.BlockEntityData{
			Type:   syncedEntity.GetType(),
			Synced: &syncedEntity,
		}
		return &data
	} else {
		nohyphenId := strings.ReplaceAll(syncedBlock.SyncedFrom.BlockId, "-", "")
		syncedEntity := domain.SyncedEntity(nohyphenId)
		data := domain.BlockEntityData{
			Type:   syncedEntity.GetType(),
			Synced: &syncedEntity,
		}
		return &data
	}
}

func makeChildDatabaseData(database_id string) (*domain.BlockEntityData, error) {
	dbData, err := GetDBItem(database_id)
	if err != nil {
		if errors.Is(err, domain.ErrNotionErrorResponse) {
			return nil, domain.ErrNotionErrorResponse
		}
		fmt.Println("error in usecase/makeChildDatabaseData/GetDBItem")
		return nil, err
	}
	dbQueryData, err := GetChildDB(database_id)
	if err != nil {
		fmt.Println("error in usecase/makeChildDatabaseData/GetDBQuery")
		return nil, err
	}
	childdbEntity := domain.ChildDBEntity{
		DatabaseData: dbData,
		QueryData:    dbQueryData,
	}
	data := domain.BlockEntityData{
		Type:    childdbEntity.GetType(),
		ChildDB: &childdbEntity,
	}
	return &data, nil
}
