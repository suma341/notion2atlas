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
func (c InfoEntity) ToPageEntity() PageEntity {
	return PageEntity{
		Id:           c.Id,
		CurriculumId: c.Id,
		IconType:     c.IconType,
		IconUrl:      c.IconUrl,
		ParentId:     "",
		CoverUrl:     c.CoverUrl,
		CoverType:    c.CoverType,
		Order:        c.Order,
		Title:        c.Title,
	}
}
func (c InfoEntity) GetTitle() string {
	return c.Title
}

func (inf InfoEntity) CompareQueryEntityTime(q2 DBQueryEntity) (bool, error) {
	t1, err := inf.GetTime()
	if err != nil {
		fmt.Println("error in utils/CompareQueryEntityTime/inf.GetTime")
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
