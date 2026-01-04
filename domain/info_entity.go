package domain

import (
	"fmt"
	"time"
)

type InfoEntity struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Order          int    `json:"order"`
	IconType       string `json:"iconType"`
	IconUrl        string `json:"iconUrl"`
	CoverType      string `json:"coverType"`
	CoverUrl       string `json:"coverUrl"`
	LastEditedTime string `json:"last_edited_time"`
	Update         bool   `json:"update"`
}

func (c InfoEntity) GetCategories() []string {
	return []string{"部活情報"}
}
func (c InfoEntity) GetVisilities() []string {
	return []string{"基礎班", "発展班"}
}
func (c InfoEntity) GetTags() []string {
	return []string{"情報"}
}

func (c InfoEntity) GetId() string {
	return c.Id
}

func (c InfoEntity) GetLastEditedTime() string {
	return c.LastEditedTime
}

func (c InfoEntity) GetUpdate() bool {
	return c.Update
}
func (c InfoEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("error in entity/InfoEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}
func (c InfoEntity) ToPageEntity() (*PageEntity, error) {
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
		"info",
		c.LastEditedTime,
	)
	if err != nil {
		fmt.Println("error in domain/CategoryEntity.ToPageEntity/NewPageEntity")
		return nil, err
	}
	return page, nil
}

func (c InfoEntity) GetTitle() string {
	return c.Title
}

func (inf InfoEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
	t1, err := inf.GetTime()
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
