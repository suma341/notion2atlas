package domain

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type NtPageEntity struct {
	Id             string `json:"id"`
	IconUrl        string `json:"iconUrl"`
	IconType       string `json:"iconType"`
	CoverUrl       string `json:"coverUrl"`
	CoverType      string `json:"coverType"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	LastEditedTime string `json:"last_edited_time"`
	InTrash        bool   `json:"in_trash"`
}

func (c NtPageEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("❌ error in entity/CurriculumEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}

func (curr NtPageEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
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

func (p NtPageEntity) GetTitle() string {
	return p.Title
}
func (p NtPageEntity) GetId() string {
	return p.Id
}
func (p NtPageEntity) GetIcon() (IconType string, IconUrl string) {
	return p.IconType, p.IconUrl
}
func (p NtPageEntity) GetCover() (CoverType string, CoverUrl string) {
	return p.CoverType, p.CoverUrl
}

func NewNtPageEntity(
	Id string,
	IconUrl string,
	IconType string,
	CoverUrl string,
	CoverType string,
	Title string,
	Type string,
	LastEditedTime string,
	InTrash bool,
) NtPageEntity {
	return NtPageEntity{
		Id:             Id,
		IconUrl:        IconUrl,
		IconType:       IconType,
		CoverUrl:       CoverUrl,
		CoverType:      CoverType,
		Title:          Title,
		Type:           Type,
		LastEditedTime: LastEditedTime,
		InTrash:        InTrash,
	}
}

func ResNtPageEntity(res map[string]any, type_ string) (*NtPageEntity, error) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("❌ error in domain/Res2NtPageEntity/json.Marshal(res)")
		return nil, err
	}
	var obj NtObjectProbe
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println("❌ error in domain/ResNtPageEntity/json.Marshal(jsonBytes, &obj)")
		return nil, err
	}
	if obj.Object != "error" {
		var block PageProperty
		err = json.Unmarshal(jsonBytes, &block)
		if err != nil {
			fmt.Println("❌ error in domain/ResNtPageEntity/json.Marshal(jsonBytes, &block)")
			return nil, err
		}
		id := strings.ReplaceAll(block.Id, "-", "")
		title := block.Properties.Title.GetCombinedPlainText()
		var iconUrl = ""
		var iconType = ""
		if block.Icon != nil {
			iconUrl = block.Icon.GetIconUrl()
			iconType = block.Icon.Type
		}
		var coverUrl = ""
		coverType := ""
		if block.Cover != nil {
			coverType = block.Cover.Type
			coverUrl = block.Cover.GetCoverUrl()
		}
		var page = NewNtPageEntity(
			id,
			iconUrl,
			iconType,
			coverUrl,
			coverType,
			title,
			type_,
			block.LastEditedTime,
			block.InTrash,
		)
		return &page, nil
	} else {
		return nil, ErrNotionErrorResponse
	}
}
