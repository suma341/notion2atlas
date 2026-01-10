package fileUC

import "fmt"

type ChangeItem struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"` // delete | add | update
}

func (c ChangeItem) GetId() string {
	return c.Id
}

func NewChangeItem(id string, title string, type_ string) (*ChangeItem, error) {
	if type_ != "delete" && type_ != "add" && type_ != "update" {
		return nil, fmt.Errorf("unexpect type of changeItem: %s", type_)
	}
	return &ChangeItem{
		Id:    id,
		Title: title,
		Type:  type_,
	}, nil
}
