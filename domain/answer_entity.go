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

func (c AnswerEntity) GetCategories() []string {
	return []string{"解答"}
}
func (c AnswerEntity) GetVisilities() []string {
	return []string{"基礎班", "発展班"}
}
func (c AnswerEntity) GetTags() []string {
	return []string{"解答"}
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
func (c AnswerEntity) ToPageEntity() (*PageEntity, error) {
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
		"answer",
		c.LastEditedTime,
	)
	if err != nil {
		fmt.Println("error in domain/CategoryEntity.ToPageEntity/NewPageEntity")
		return nil, err
	}
	return page, nil
}

func (c AnswerEntity) GetTitle() string {
	return c.Title
}

func (ans AnswerEntity) CompareQueryEntityTime(q2 NtBlock) (bool, error) {
	t1, err := ans.GetTime()
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
