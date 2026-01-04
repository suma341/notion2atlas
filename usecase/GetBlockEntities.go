package usecase

import (
	"errors"
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/usecase/fileUC"
)

func GetBlockEntities(
	res domain.NTBlockEntity,
	buffer []domain.BlockEntity,
	curriculumId string,
	pageId string,
	i int,
	type_ string,
) ([]domain.BlockEntity, error) {

	switch res.Type {
	case "paragraph":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Paragraph, type_)
	case "quote":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Quote, type_)
	case "toggle":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Toggle, type_)
	case "bulleted_list_item":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.BulletedListItem, type_)
	case "numbered_list_item":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.NumberedListItem, type_)
	case "to_do":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.ToDo, type_)
	case "heading_1":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Heading1, type_)
	case "heading_2":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Heading2, type_)
	case "heading_3":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Heading3, type_)
	case "image":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Image, type_)
	case "video":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Video, type_)
	case "embed":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Embed, type_)
	case "bookmark":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Bookmark, type_)
	case "table":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Table, type_)
	case "table_row":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.TableRow, type_)
	case "child_page":
		return appendBlock(buffer, res, curriculumId, pageId, i, 0, type_)
	case "link_to_page":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.LinkToPage, type_)
	case "code":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Code, type_)
	case "synced_block":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.SyncedBlock, type_)
	case "child_database":
		return appendBlock(buffer, res, curriculumId, pageId, i, 0, type_)
	case "callout":
		return appendBlock(buffer, res, curriculumId, pageId, i, res.Callout, type_)
	default:
		return appendBlock(buffer, res, curriculumId, pageId, i, 0, type_)
	}
}

func appendBlock(
	buffer []domain.BlockEntity,
	res domain.NTBlockEntity,
	curriculumId string,
	pageId string,
	i int,
	block any,
	// pageBuffer []domain.PageEntity,
	type_ string,
) ([]domain.BlockEntity, error) {
	var data *domain.BlockEntityData
	var err error = nil
	parentId, err := res.GetParentId()
	if err != nil {
		fmt.Println("error in usecase/appendBlock/res.GetParentId")
		return buffer, err
	}
	switch res.Type {
	case "paragraph", "quote", "toggle", "bulleted_list_item", "numbered_list_item":
		obj, ok := block.(*domain.ParagraphProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.ParagraphProperty)")
		}
		data, err = makeParagraphData(*obj, type_)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeParagraphData")
			return buffer, err
		}

	case "to_do":
		obj, ok := block.(*domain.ToDoProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.ParagraphProperty)")
		}
		data, err = makeToDoData(*obj, type_)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeParagraphData")
			return buffer, err
		}
	case "heading_1", "heading_2", "heading_3":
		obj, ok := block.(*domain.HeaderProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.Header)")
		}
		data, err = makeHeaderData(*obj, type_)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeHeaderData")
			return buffer, err
		}
	case "image", "video":
		obj, ok := block.(*domain.ImageProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.ImageProperty)")
		}
		data, err = makeImageData(*obj, res.Id, pageId, type_)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeImageData")
			return buffer, err
		}
	case "embed", "bookmark":
		obj, ok := block.(*domain.EmbedProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.EmbedProperty)")
		}
		data, err = makeEmbedData(*obj, res.Type)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeEmbedData")
			return buffer, err
		}
	case "table":
		obj, ok := block.(*domain.TableProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.TableProperty)")
		}
		data, err = makeTableData(*obj)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeTableData")
			return buffer, err
		}
	case "table_row":
		obj, ok := block.(*domain.TableRowProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.TableRowProperty)")
		}
		data, err = makeTableRowData(*obj, type_)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeTableRowData")
			return buffer, err
		}
	case "link_to_page":
		obj, ok := block.(*domain.LinkToPageProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.LinkToPageProperty)")
		}
		data, err = makeLinkToPageData(*obj, type_)
		if err != nil {
			if errors.Is(err, domain.ErrNotionErrorResponse) {
				return buffer, nil
			}
			fmt.Println("error in usecase/appendBlock/makeLinkToPageData")
			return buffer, err
		}
	case "code":
		obj, ok := block.(*domain.CodeProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.CodeProperty)")
		}
		data, err = makeCodeData(*obj, type_)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeCodeData")
			return buffer, err
		}
	case "callout":
		obj, ok := block.(*domain.CalloutProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.CodeProperty)")
		}
		data, err = makeCalloutData(*obj, type_)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/makeCodeData")
			return buffer, err
		}
	case "synced_block":
		obj, ok := block.(*domain.SyncedProperty)
		if !ok {
			return buffer, fmt.Errorf("type convert failed: block.(*domain.SyncedBlock)")
		}
		data = makeSyncedBlockData(*obj)
		if *data.Synced == "original" {
			var data_ = domain.BlockEntityData{
				Type:   "synced",
				Synced: data.Synced,
			}
			err := fileUC.UpsertSyncedFile(domain.BlockEntity{
				Id:           res.Id,
				Type:         res.Type,
				ParentId:     parentId,
				CurriculumId: curriculumId,
				PageId:       pageId,
				Data:         data_,
				Order:        i,
			})
			if err != nil {
				fmt.Println("error in usecase/appendBlock/UpsertSyncedFile")
				return buffer, err
			}
		}
	case "child_database":
		data, err = makeChildDatabaseData(res.Id)
		if err != nil {
			if errors.Is(err, domain.ErrNotionErrorResponse) {
				return buffer, nil
			}
			fmt.Println("error in usecase/appendBlock/makeChildDatabaseData")
			return buffer, err
		}
	case "child_page":
		var pageDataAddress *domain.NtPageEntity
		data, pageDataAddress, err = makeChildPageData(res.Id, type_)
		if err != nil {
			if errors.Is(err, domain.ErrNotionErrorResponse) {
				return buffer, nil
			}
			fmt.Println("error in usecase/appendBlock/makeChildPageData")
			return buffer, err
		}
		if pageDataAddress == nil {
			return buffer, fmt.Errorf("pageDataAddress is nil")
		}
		pageEntity, err := domain.NewPageEntity(
			res.Id,
			curriculumId,
			pageDataAddress.IconType,
			pageDataAddress.IconUrl,
			pageDataAddress.CoverUrl,
			pageDataAddress.CoverType,
			i,
			parentId,
			pageDataAddress.Title,
			type_,
			pageDataAddress.LastEditedTime,
		)
		if err != nil {
			fmt.Println("error in usecase/appendBlock/domain.NewPageEntity")
			return buffer, err
		}
		filemanager.WriteJson(pageEntity, fmt.Sprintf("%s%s_new.json", constants.TMP_DIR, res.Id))
	default:
		data = &domain.BlockEntityData{
			Type: "none",
		}
	}
	domain := domain.BlockEntity{
		Id:           res.Id,
		Type:         res.Type,
		ParentId:     parentId,
		CurriculumId: curriculumId,
		PageId:       pageId,
		Data:         *data,
		Order:        i,
	}
	return append(buffer, domain), nil
}
