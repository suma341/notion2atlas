package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NtPageEntity struct {
	Id        string `json:"id"`
	IconUrl   string `json:"iconUrl"`
	IconType  string `json:"iconType"`
	CoverUrl  string `json:"coverUrl"`
	CoverType string `json:"coverType"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Object    string `json:"object"` //
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

func ResNtPageEntity(res map[string]any, type_ string) (*NtPageEntity, error) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error in domain/Res2NtPageEntity/json.Marshal(res)")
		return nil, err
	}
	var obj NtObjectProbe
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println("error in domain/ResNtPageEntity/json.Marshal(jsonBytes, &obj)")
		return nil, err
	}
	if obj.Object != "error" {
		var block PageProperty
		err = json.Unmarshal(jsonBytes, &block)
		if err != nil {
			fmt.Println("error in domain/ResNtPageEntity/json.Marshal(jsonBytes, &block)")
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
		var page = NtPageEntity{
			Id:        id,
			IconUrl:   iconUrl,
			IconType:  iconType,
			CoverUrl:  coverUrl,
			CoverType: coverType,
			Title:     title,
			Type:      type_,
		}
		return &page, nil
	} else {
		return nil, ErrNotionErrorResponse
	}
}
