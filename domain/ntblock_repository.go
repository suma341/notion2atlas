package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NTBlockRepository struct {
	Id          string     `json:"id"`
	Type        string     `json:"type"`
	HasChildren bool       `json:"has_children"`
	Parent      ParentData `json:"parent"`
	// types
	Paragraph        *ParagraphProperty  `json:"paragraph,omitempty"`
	Quote            *ParagraphProperty  `json:"quote,omitempty"`
	Toggle           *ParagraphProperty  `json:"toggle,omitempty"`
	BulletedListItem *ParagraphProperty  `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ParagraphProperty  `json:"numbered_list_item,omitempty"`
	ToDo             *ParagraphProperty  `json:"ToDo,omitempty"`
	Heading1         *HeaderProperty     `json:"heading_1,omitempty"`
	Heading2         *HeaderProperty     `json:"heading_2,omitempty"`
	Heading3         *HeaderProperty     `json:"heading_3,omitempty"`
	Callout          *CalloutProperty    `json:"callout,omitempty"`
	Image            *ImageProperty      `json:"image,omitempty"`
	Video            *ImageProperty      `json:"video,omitempty"`
	Embed            *EmbedProperty      `json:"embed,omitempty"`
	Bookmark         *EmbedProperty      `json:"bookmark,omitempty"`
	Table            *TableProperty      `json:"table,omitempty"`
	TableRow         *TableRowProperty   `json:"table_row,omitempty"`
	LinkToPage       *LinkToPageProperty `json:"link_to_page,omitempty"`
	Code             *CodeProperty       `json:"code,omitempty"`
	SyncedBlock      *SyncedProperty     `json:"synced_block,omitempty"`
}

type ParentData struct {
	PageId     *string `json:"page_id,omitempty"`
	BlockId    *string `json:"block_id,omitempty"`
	DatabaseId *string `json:"database_id,omitempty"`
	Type       string  `json:"type"`
}

func (b NTBlockRepository) GetParentId() (string, error) {
	switch b.Parent.Type {
	case "page_id":
		return strings.ReplaceAll(*b.Parent.PageId, "-", ""), nil
	case "block_id":
		return strings.ReplaceAll(*b.Parent.BlockId, "-", ""), nil
	case "database_id":
		return strings.ReplaceAll(*b.Parent.DatabaseId, "-", ""), nil
	}
	return "", fmt.Errorf("unexpected parent type")
}

func Res2NTBlockRepository(res map[string]any) (*NTBlockRepository, error) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error in domain/Res2NTBlockRepository/json.Marshal(res)")
		return nil, err
	}
	var obj NTBlockRepository
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println("error in domain/Res2NTBlockRepository/son.Unmarshal(jsonBytes, &obj)")
		return nil, err
	}
	obj.Id = strings.ReplaceAll(obj.Id, "-", "")
	return &obj, nil
}
