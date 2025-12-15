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
	Category       string   `json:"category"`
	IconType       string   `json:"iconType"`
	IconUrl        string   `json:"iconUrl"`
	CoverType      string   `json:"coverType"`
	CoverUrl       string   `json:"coverUrl"`
	LastEditedTime string   `json:"last_edited_time"`
	Update         bool     `json:"update"`
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

func (curr CurriculumEntity) CompareQueryEntityTime(q2 DBQueryEntity) (bool, error) {
	t1, err := curr.GetTime()
	if err != nil {
		fmt.Println("error in utils/CompareQueryEntityTime/curr.GetTime")
		return false, err
	}
	if t1 == nil {
		return false, fmt.Errorf("unexpected: t1 is nil")
	}
	t2, err := q2.GetTime()
	if err != nil {
		fmt.Println("error in utils/CompareQueryEntityTime/q2.GetTime")
		return false, err
	}
	if t2 == nil {
		return false, fmt.Errorf("unexpected: t2 is nil")
	}
	isEqual := t1.Equal(*t2)
	return isEqual, nil
}
