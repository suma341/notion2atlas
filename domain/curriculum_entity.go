package domain

import (
	"fmt"
	"time"
)

type CurriculumEntity struct {
	Id             string   `json:"id"`
	Title          string   `json:"title"`
	Tag            []string `json:"tag"`
	Visibility     []string `json:"visibility"`
	Order          int      `json:"order"`
	Category       []string `json:"category"`
	IconType       string   `json:"iconType"`
	IconUrl        string   `json:"iconUrl"`
	CoverType      string   `json:"coverType"`
	CoverUrl       string   `json:"coverUrl"`
	LastEditedTime string   `json:"last_edited_time"`
	Update         bool     `json:"update"`
}

func (c CurriculumEntity) GetCategories() []string {
	return c.Category
}
func (c CurriculumEntity) GetVisilities() []string {
	return c.Visibility
}
func (c CurriculumEntity) GetTags() []string {
	return c.Tag
}

func (c CurriculumEntity) GetId() string {
	return c.Id
}

func (c CurriculumEntity) GetLastEditedTime() string {
	return c.LastEditedTime
}
func (c CurriculumEntity) GetUpdate() bool {
	return c.Update
}

func (c CurriculumEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("error in entity/CurriculumEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}
func (c CurriculumEntity) ToPageEntity() (*PageEntity, error) {
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
		fmt.Println("error in domain/CurriculumEntity.ToPageEntity/NewPageEntity")
		return nil, err
	}
	return page, nil
}
func (c CurriculumEntity) GetTitle() string {
	return c.Title
}

func (curr CurriculumEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
	t1, err := curr.GetTime()
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
