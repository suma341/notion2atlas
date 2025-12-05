package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func GetBlockEntities(
	res domain.NTBlockRepository,
	buffer []domain.BlockEntity,
	curriculumId string,
	pageId string,
	i int,
	pageBuffer []domain.PageEntity,
) ([]domain.BlockEntity, []domain.PageEntity, error) {

	switch res.Type {
	case "paragraph":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Paragraph, pageBuffer)
	case "quote":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Quote, pageBuffer)
	case "toggle":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Toggle, pageBuffer)
	case "bulleted_list_item":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.BulletedListItem, pageBuffer)
	case "numbered_list_item":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.NumberedListItem, pageBuffer)
	case "to_do":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.ToDo, pageBuffer)
	case "heading_1":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Heading1, pageBuffer)
	case "heading_2":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Heading2, pageBuffer)
	case "heading_3":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Heading3, pageBuffer)
	case "image":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Image, pageBuffer)
	case "video":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Video, pageBuffer)
	case "embed":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Embed, pageBuffer)
	case "bookmark":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Bookmark, pageBuffer)
	case "table":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Table, pageBuffer)
	case "table_row":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.TableRow, pageBuffer)
	case "child_page":
		return appendBlock(buffer, res, curriculumId, pageId, i, 0, pageBuffer)
	case "link_to_page":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.LinkToPage, pageBuffer)
	case "code":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Code, pageBuffer)
	case "synced_block":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.SyncedBlock, pageBuffer)
	case "child_database":
		return appendBlock(buffer, res, curriculumId, pageId, i, 0, pageBuffer)
	case "callout":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Callout, pageBuffer)
	default:
		return appendBlock(buffer, res, curriculumId, pageId, i, 0, pageBuffer)
	}
}

func appendBlock(
	buffer []domain.BlockEntity,
	res domain.NTBlockRepository,
	curriculumId string,
	pageId string,
	i int,
	block any,
	pageBuffer []domain.PageEntity,
) ([]domain.BlockEntity, []domain.PageEntity, error) {
	var data any
	var err error = nil
	parentId, err := res.GetParentId()
	if err != nil {
		fmt.Println("error in usecase/appendBlock/res.GetParentId")
		return buffer, pageBuffer, err
	}
	switch res.Type {
	case "paragraph", "quote", "toggle", "bulleted_list_item", "numbered_list_item", "to_do":
		obj, ok := block.(*domain.ParagraphProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.ParagraphProperty)")
		}
		data, err = makeParagraphData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeParagraphData")
			return buffer, pageBuffer, err
		}
	case "heading_1", "heading_2", "heading_3":
		obj, ok := block.(*domain.HeaderProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.Header)")
		}
		data, err = makeHeaderData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeHeaderData")
			return buffer, pageBuffer, err
		}
	case "image", "video":
		obj, ok := block.(*domain.ImageProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.ImageProperty)")
		}
		data, err = makeImageData(*obj, res.Id, pageId)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeImageData")
			return buffer, pageBuffer, err
		}
	case "embed", "bookmark":
		obj, ok := block.(*domain.EmbedProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.EmbedProperty)")
		}
		data, err = makeEmbedData(*obj, res.Type)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeEmbedData")
			return buffer, pageBuffer, err
		}
	case "table":
		obj, ok := block.(*domain.TableProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.TableProperty)")
		}
		data, err = makeTableData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeTableData")
			return buffer, pageBuffer, err
		}
	case "table_row":
		obj, ok := block.(*domain.TableRowProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.TableRowProperty)")
		}
		data, err = makeTableRowData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeTableRowData")
			return buffer, pageBuffer, err
		}
	case "link_to_page":
		obj, ok := block.(*domain.LinkToPageProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.LinkToPageProperty)")
		}
		data, err = makeLinkToPageData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeLinkToPageData")
			return buffer, pageBuffer, err
		}
	case "code":
		obj, ok := block.(*domain.CodeProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.CodeProperty)")
		}
		data, err = makeCodeData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeCodeData")
			return buffer, pageBuffer, err
		}
	case "callout":
		obj, ok := block.(*domain.CalloutProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.CodeProperty)")
		}
		data, err = makeCalloutData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeCodeData")
			return buffer, pageBuffer, err
		}
	case "synced_block":
		obj, ok := block.(*domain.SyncedProperty)
		if !ok {
			return buffer, pageBuffer, fmt.Errorf("type convert failed: block.(*domain.SyncedBlock)")
		}
		data = makeSyncedBlockData(*obj)
		if data == "original" {
			err := UpsertSyncedFile(domain.BlockEntity{
				Id:           res.Id,
				Type:         res.Type,
				ParentId:     parentId,
				CurriculumId: curriculumId,
				PageId:       pageId,
				Data:         data,
				Order:        i,
			})
			if err != nil {
				fmt.Println("error in usecase/appendBlock/UpsertSyncedFile")
				return buffer, pageBuffer, err
			}
		}
	case "child_database":
		data, err = makeChildDatabaseData(res.Id)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeChildDatabaseData")
			return buffer, pageBuffer, err
		}
	case "child_page":
		var pageDataAddress *domain.NTPageRepository
		data, pageDataAddress, err = makeChildPageData(res.Id)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeChildPageData")
			return buffer, pageBuffer, err
		}
		if pageDataAddress == nil {
			return buffer, pageBuffer, fmt.Errorf("pageDataAddress is nil")
		}
		pageEntity := domain.PageEntity{
			Id:           res.Id,
			CurriculumId: curriculumId,
			IconType:     pageDataAddress.IconType,
			IconUrl:      pageDataAddress.IconUrl,
			ParentId:     parentId,
			Order:        i,
			CoverUrl:     pageDataAddress.CoverUrl,
			CoverType:    pageDataAddress.CoverType,
			Title:        pageDataAddress.Title,
		}
		pageBuffer = append(pageBuffer, pageEntity)
	default:
		data = "_"
	}
	domain := domain.BlockEntity{
		Id:           res.Id,
		Type:         res.Type,
		ParentId:     parentId,
		CurriculumId: curriculumId,
		PageId:       pageId,
		Data:         data,
		Order:        i,
	}
	return append(buffer, domain), pageBuffer, nil
}
