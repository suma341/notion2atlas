package domain

import (
	"fmt"
	"time"
)

type PageEntity struct {
	Id             string `json:"id"`
	CurriculumId   string `json:"curriculumId"`
	IconType       string `json:"iconType"`
	IconUrl        string `json:"iconUrl"`
	CoverUrl       string `json:"coverUrl"`
	CoverType      string `json:"coverType"`
	Order          int    `json:"order"`
	ParentId       string `json:"parentId"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	LastEditedTime string `json:"last_edited_time"`
}

func (c PageEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("❌ error in entity/CurriculumEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}

func (curr PageEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
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
func (p PageEntity) GetTitle() string {
	return p.Title
}
func (p PageEntity) GetId() string {
	return p.Id
}
func (p PageEntity) GetIcon() (IconType string, IconUrl string) {
	return p.IconType, p.IconUrl
}
func (p PageEntity) GetCover() (CoverType string, CoverUrl string) {
	return p.CoverType, p.CoverUrl
}
func (p PageEntity) ChangePageEntityUrl(iconUrl string, coverUrl string) (*PageEntity, error) {
	entity, err := NewPageEntity(
		p.Id,
		p.CurriculumId,
		p.IconType,
		iconUrl,
		coverUrl,
		p.CoverType,
		p.Order,
		p.ParentId,
		p.Title,
		p.Type,
		p.LastEditedTime,
	)
	if err != nil {
		fmt.Println("❌ error in domain/PageEntity.ChangePageEntityUrl/NewPageEntity")
		return nil, err
	}
	return entity, nil
}

func NewPageEntity(
	Id string,
	CurriculumId string,
	IconType string,
	IconUrl string,
	CoverUrl string,
	CoverType string,
	Order int,
	ParentId string,
	Title string,
	Type string,
	LastEditedTime string,
) (*PageEntity, error) {
	if Type != "curriculum" && Type != "info" && Type != "answer" {
		return nil, fmt.Errorf("unexpected type: %s", Type)
	}
	return &PageEntity{
		Id:             Id,
		CurriculumId:   CurriculumId,
		IconType:       IconType,
		IconUrl:        IconUrl,
		CoverUrl:       CoverUrl,
		CoverType:      CoverType,
		Order:          Order,
		ParentId:       ParentId,
		Title:          Title,
		Type:           Type,
		LastEditedTime: LastEditedTime,
	}, nil
}

func CreatePage(title string, iconType string, iconUrl string, id string) PageEntity {
	return PageEntity{
		Id:           id,
		CurriculumId: "",
		IconType:     iconType,
		IconUrl:      iconUrl,
		CoverUrl:     "",
		CoverType:    "",
		Order:        0,
		ParentId:     "",
		Title:        title,
		Type:         "",
	}
}
