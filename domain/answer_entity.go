package domain

import (
	"fmt"
	"time"
)

type AnswerEntity struct {
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

func (c AnswerEntity) GetId() string {
	return c.Id
}

func (c AnswerEntity) GetLastEditedTime() string {
	return c.LastEditedTime
}

func (c AnswerEntity) GetUpdate() bool {
	return c.Update
}

func (c AnswerEntity) GetTime() (*time.Time, error) {
	lastEditedTime, err := time.Parse(time.RFC3339, c.LastEditedTime)
	if err != nil {
		fmt.Println("error in entity/AnswerEntity/GetTime")
		return nil, err
	}
	return &lastEditedTime, nil
}
func (c AnswerEntity) ToPageEntity() PageEntity {
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
func (c AnswerEntity) GetTitle() string {
	return c.Title
}

func (ans AnswerEntity) CompareQueryEntityTime(q2 DBQueryEntity) (bool, error) {
	t1, err := ans.GetTime()
	if err != nil {
		fmt.Println("error in utils/CompareQueryEntityTime/ans.GetTime")
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
