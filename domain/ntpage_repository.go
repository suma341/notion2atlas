package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NTPageRepository struct {
	Id        string `json:"id"`
	IconUrl   string `json:"iconUrl"`
	IconType  string `json:"iconType"`
	CoverUrl  string `json:"coverUrl"`
	CoverType string `json:"coverType"`
	Title     string `json:"title"`
	Type      string `json:"type"`
}

func (p NTPageRepository) GetTitle() string {
	return p.Title
}
func (p NTPageRepository) GetId() string {
	return p.Id
}
func (p NTPageRepository) GetIcon() (IconType string, IconUrl string) {
	return p.IconType, p.IconUrl
}
func (p NTPageRepository) GetCover() (CoverType string, CoverUrl string) {
	return p.CoverType, p.CoverUrl
}

func ResNTPageRepository(res map[string]any, type_ string) (*NTPageRepository, error) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error in domain/Res2NTPageRepository/json.Marshal(res)")
		return nil, err
	}
	var obj PageProperty
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println("error in domain/ResNTPageRepository/json.Marshal(jsonBytes, &obj)")
		return nil, err
	}
	id := strings.ReplaceAll(obj.Id, "-", "")
	title := obj.Properties.Title.GetCombinedPlainText()
	var iconUrl = ""
	var iconType = ""
	if obj.Icon != nil {
		iconUrl = obj.Icon.GetIconUrl()
		iconType = obj.Icon.Type
	}
	var coverUrl = ""
	coverType := ""
	if obj.Cover != nil {
		coverType = obj.Cover.Type
		coverUrl = obj.Cover.GetCoverUrl()
	}
	var page = NTPageRepository{
		Id:        id,
		IconUrl:   iconUrl,
		IconType:  iconType,
		CoverUrl:  coverUrl,
		CoverType: coverType,
		Title:     title,
		Type:      type_,
	}
	return &page, nil
}
