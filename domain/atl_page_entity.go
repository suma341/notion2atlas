package domain

import (
	"fmt"
	"time"
)

type AtlPageEntity struct {
	Id             string   `json:"id"`
	CurriculumId   string   `json:"curriculumId"`
	IconType       string   `json:"iconType"`
	IconUrl        string   `json:"iconUrl"`
	CoverUrl       string   `json:"coverUrl"`
	CoverType      string   `json:"coverType"`
	Order          int      `json:"order"`
	ParentId       string   `json:"parentId"`
	Title          string   `json:"title"`
	Type           string   `json:"type"`
	Ogp            PageOgp  `json:"ogp"`
	Visibility     []string `json:"visibility"`
	Tag            []string `json:"tag"`
	Category       []string `json:"category"`
	LastEditedTime string   `json:"last_edited_time"`
}

func (c AtlPageEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("❌ error in entity/CurriculumEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}

func (curr AtlPageEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
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

func (p AtlPageEntity) GetId() string {
	return p.Id
}

func (p AtlPageEntity) GetIcon() (IconType string, IconUrl string) {
	return p.IconType, p.IconUrl
}
func (p AtlPageEntity) GetCover() (CoverType string, CoverUrl string) {
	return p.CoverType, p.CoverUrl
}
func (p AtlPageEntity) GetTitle() string {
	return p.Title
}

func (p AtlPageEntity) ToPageEntity() (*PageEntity, error) {
	return NewPageEntity(
		p.Id,
		p.CurriculumId,
		p.IconType,
		p.IconUrl,
		p.CoverUrl,
		p.CoverType,
		p.Order,
		p.ParentId,
		p.Title,
		p.Type,
		p.LastEditedTime,
	)
}

func (p AtlPageEntity) ChangePageEntityUrl(iconUrl string, coverUrl string) AtlPageEntity {
	entity := NewAtlPageEntity(
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
		p.Ogp,
		p.Visibility,
		p.Tag,
		p.Category,
		p.LastEditedTime,
	)
	return entity
}

func NewAtlPageEntity(
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
	Ogp PageOgp,
	Visibility []string,
	Tag []string,
	Category []string,
	LastEditedTime string,
) AtlPageEntity {
	return AtlPageEntity{
		Id:             Id,
		CurriculumId:   CurriculumId,
		IconType:       IconType,
		IconUrl:        IconUrl,
		CoverType:      CoverType,
		CoverUrl:       CoverUrl,
		Order:          Order,
		ParentId:       ParentId,
		Title:          Title,
		Type:           Type,
		Ogp:            Ogp,
		Visibility:     Visibility,
		Tag:            Tag,
		Category:       Category,
		LastEditedTime: LastEditedTime,
	}
}

type PageOgp struct {
	FirstText string `json:"first_text"`
	ImagePath string `json:"image_path"`
}
