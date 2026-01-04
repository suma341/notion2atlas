package domain

import (
	"fmt"
	"time"
)

type CategoryEntity struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	IconType          string `json:"iconType"`
	IconUrl           string `json:"iconUrl"`
	CoverType         string `json:"coverType"`
	CoverUrl          string `json:"coverUrl"`
	IsBasicCurriculum bool   `json:"is_basic_curriculum"`
	Order             int    `json:"order"`
	LastEditedTime    string `json:"last_edited_time"`
	Update            bool   `json:"update"`
}

func (c CategoryEntity) GetId() string {
	return c.Id
}

func (c CategoryEntity) GetLastEditedTime() string {
	return c.LastEditedTime
}
func (c CategoryEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("error in entity/CategoryEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}

func (c CategoryEntity) GetUpdate() bool {
	return c.Update
}

func (c CategoryEntity) ToPageEntity() (*PageEntity, error) {
	page, err := NewPageEntity(
		c.Id,
		c.Id,
		c.IconType,
		c.IconUrl,
		c.CoverUrl,
		c.CoverType,
		c.Order,
		"",
		c.Title,
		"curriculum",
		c.LastEditedTime,
	)
	if err != nil {
		fmt.Println("error in domain/CategoryEntity.ToPageEntity/NewPageEntity")
		return nil, err
	}
	return page, nil
}
func (c CategoryEntity) GetTitle() string {
	return c.Title
}

func (cat CategoryEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
	t1, err := cat.GetTime()
	if err != nil {
		return false, err
	}
	t2, err := q2.GetTime()
	if err != nil {
		return false, err
	}

	t1u := t1.UTC().Truncate(time.Second)
	t2u := t2.UTC().Truncate(time.Second)

	return t1u.Equal(t2u), nil
}

func CreateStaticCategory(
	Id string,
	Title string,
	IconType string,
	IconUrl string,

) CategoryEntity {
	return CategoryEntity{
		Id:                Id,
		Title:             Title,
		Description:       "",
		IconType:          IconType,
		IconUrl:           IconUrl,
		IsBasicCurriculum: false,
		Order:             1,
	}
}
