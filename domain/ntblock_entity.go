package domain

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type NTBlockEntity struct {
	Id             string     `json:"id"`
	Type           string     `json:"type"`
	HasChildren    bool       `json:"has_children"`
	Parent         ParentData `json:"parent"`
	Object         string     `json:"object"` //block or list
	LastEditedTime string     `json:"last_edited_time"`
	// types
	Paragraph        *ParagraphProperty  `json:"paragraph,omitempty"`
	Quote            *ParagraphProperty  `json:"quote,omitempty"`
	Toggle           *ParagraphProperty  `json:"toggle,omitempty"`
	BulletedListItem *ParagraphProperty  `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ParagraphProperty  `json:"numbered_list_item,omitempty"`
	ToDo             *ToDoProperty       `json:"to_do,omitempty"`
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

func (c NTBlockEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("❌ error in entity/CurriculumEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}

func (curr NTBlockEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
	t1, err := curr.GetTime()
	if err != nil {
		fmt.Println("❌ error in utils/CompareQueryEntityTime/curr.GetTime")
		return false, err
	}
	if t1 == nil {
		return false, fmt.Errorf("unexpected: t1 is nil")
	}
	t2, err := q2.GetTime()
	if err != nil {
		fmt.Println("❌ error in utils/CompareQueryEntityTime/q2.GetTime")
		return false, err
	}
	if t2 == nil {
		return false, fmt.Errorf("unexpected: t2 is nil")
	}
	isEqual := t1.Equal(*t2)
	return isEqual, nil
}

type ParentData struct {
	PageId     *string `json:"page_id,omitempty"`
	BlockId    *string `json:"block_id,omitempty"`
	DatabaseId *string `json:"database_id,omitempty"`
	Type       string  `json:"type"`
}

func (b NTBlockEntity) GetParentId() (string, error) {
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

func Res2NTBlockEntity(res map[string]any) (*NTBlockEntity, error) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("❌ error in domain/Res2NTBlockEntity/json.Marshal(res)")
		return nil, err
	}
	var obj NtObjectProbe
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println("❌ error in domain/Res2NTBlockEntity/son.Unmarshal(jsonBytes, &obj)")
		return nil, err
	}
	switch obj.Object {
	case "block":
		var block NTBlockEntity
		if err := json.Unmarshal(jsonBytes, &block); err != nil {
			return nil, err
		}
		block.Id = strings.ReplaceAll(block.Id, "-", "")
		return &block, nil
	case "error":
		var ntErr NTErrorEntity
		if err := json.Unmarshal(jsonBytes, &ntErr); err != nil {
			return nil, err
		}
		return nil, ErrNotionErrorResponse
	}
	return nil, fmt.Errorf("unexpected: object type is Neither block or error")
}
