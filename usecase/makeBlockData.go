package usecase

import (
	"errors"
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/utils"
)

func makeParagraphData(para domain.ParagraphProperty, type_ string) (map[string]any, error) {
	richTextModels, err := ProcessRichText(para.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeParagraphData/ProcessRichText")
		return nil, err
	}
	data := map[string]any{
		"color":  para.Color,
		"parent": richTextModels,
	}
	return data, nil
}

func makeToDoData(todo domain.ToDoProperty, type_ string) (map[string]any, error) {
	richTextModels, err := ProcessRichText(todo.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeParagraphData/ProcessRichText")
		return nil, err
	}
	data := map[string]any{
		"color":   todo.Color,
		"parent":  richTextModels,
		"checked": todo.Checked,
	}
	return data, nil
}

func makeHeaderData(header domain.HeaderProperty, type_ string) (map[string]any, error) {
	richTextModels, err := ProcessRichText(header.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeHeaderData/ProcessRichText")
		return nil, err
	}
	data := map[string]any{
		"color":         header.Color,
		"parent":        richTextModels,
		"is_toggleable": header.IsToggleable,
	}
	return data, nil
}

func makeImageData(img domain.ImageProperty, blockId string, pageId string, type_ string) (map[string]any, error) {
	richTextModels, err := ProcessRichText(img.Caption, type_)
	if err != nil {
		fmt.Println("error in usecase/makeImageData/ProcessRichText")
		return nil, err
	}
	path := ""
	if img.File != nil {
		fileName, err := filemanager.DownloadFile(img.File.Url, "public/assets/"+pageId, blockId, ".png")
		if err != nil {
			fmt.Println("error in usecase/makeImageData/filemanager.DownloadFile")
			return nil, err
		}
		path = utils.GetDownloadPath(pageId, fileName)
	} else {
		fmt.Println("ℹ️ unexpected type: " + img.Type)
		filemanager.WriteJson(map[string]any{"type": img.Type}, "public/test.js")
	}
	data := map[string]any{
		"parent": richTextModels,
		"url":    path,
	}
	return data, nil
}

func makeEmbedData(embed domain.EmbedProperty, type_ string) (map[string]any, error) {
	var data map[string]any
	parent, err := ProcessRichText(embed.Caption, type_)
	if err != nil {
		fmt.Println("error in usecase/makeEmbedData/ProcessRichText")
		return nil, err
	}
	switch type_ {
	case "embed":
		canEmbed := utils.CanEmbed(embed.Url)
		data = map[string]any{
			"canEmbed": canEmbed,
			"parent":   parent,
			"url":      embed.Url,
		}
	case "bookmark":
		ogpData, err := utils.GetOGP(embed.Url)
		if err != nil {
			fmt.Println("error in usecase/makeEmbedData/utils.GetOGP")
			return nil, err
		}
		data = map[string]any{
			"parent": parent,
			"url":    embed.Url,
			"ogp":    ogpData,
		}
	}
	return data, nil
}

func makeTableData(table domain.TableProperty) (map[string]any, error) {
	data, err := domain.Struct2Map(table)
	if err != nil {
		fmt.Println("error in usecase/makeTableData/converter.Struct2Map")
		return nil, err
	}
	return data, nil
}

func makeTableRowData(table_row domain.TableRowProperty, type_ string) ([][]RichTextModel, error) {
	var cells [][]RichTextModel
	for _, cell := range table_row.Cells {
		richTextModel, err := ProcessRichText(cell, type_)
		if err != nil {
			fmt.Println("error in usecase/makeTableRowData/ProcessRichText")
			return nil, err
		}
		cells = append(cells, richTextModel)
	}
	return cells, nil
}

func makeChildPageData(pageId string, type_ string) (map[string]any, *domain.NtPageEntity, error) {
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
	data := map[string]any{
		"parent":   pageDataAddress.Title,
		"iconType": pageDataAddress.IconType,
		"iconUrl":  urls.IconUrl,
		"coverUrl": urls.CoverUrl,
	}
	pageRepo := domain.NtPageEntity{
		Id:        pageDataAddress.Id,
		IconUrl:   urls.IconUrl,
		IconType:  pageDataAddress.IconType,
		CoverUrl:  urls.CoverUrl,
		CoverType: pageDataAddress.CoverType,
		Title:     pageDataAddress.Title,
	}
	return data, &pageRepo, nil
}

func makeLinkToPageData(link_to_page domain.LinkToPageProperty, type_ string) (map[string]any, error) {
	pageId := link_to_page.PageId
	link := "/posts/curriculums/" + pageId
	pageData, err := GetPageItem(pageId, type_)
	if err != nil {
		if errors.Is(err, domain.ErrNotionErrorResponse) {
			return nil, domain.ErrNotionErrorResponse
		}
		fmt.Println("error in usecase/makeLinkToPageData/GetPageItem")
		return nil, err
	}
	urls := GetPathRewritedUrl(pageData)
	data := map[string]any{
		"link":     link,
		"iconUrl":  urls.IconUrl,
		"iconType": pageData.IconType,
		"title":    pageData.Title,
	}
	return data, nil
}

func makeCodeData(code domain.CodeProperty, type_ string) (map[string]any, error) {
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
	data := map[string]any{
		"language": code.Language,
		"caption":  caption,
		"parent":   codeContent,
	}
	return data, nil
}

func makeCalloutData(callout domain.CalloutProperty, type_ string) (map[string]any, error) {
	richText, err := ProcessRichText(callout.RichText, type_)
	if err != nil {
		fmt.Println("error in usecase/makeCalloutData/ProcessRichText")
		return nil, err
	}
	data := map[string]any{
		"icon":   callout.Icon,
		"color":  callout.Color,
		"parent": richText,
	}
	return data, nil
}

func makeSyncedBlockData(syncedBlock domain.SyncedProperty) string {
	syncedFrom := syncedBlock.SyncedFrom
	if syncedFrom == nil {
		return "original"
	} else {
		return syncedBlock.SyncedFrom.BlockId
	}
}

func makeChildDatabaseData(database_id string) (map[string]any, error) {
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
	data := map[string]any{
		"database_data": dbData,
		"query_data":    dbQueryData,
	}
	return data, nil
}
