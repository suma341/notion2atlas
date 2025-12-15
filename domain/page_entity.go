package domain

import "fmt"

type PageEntity struct {
	Id           string `json:"id"`
	CurriculumId string `json:"curriculumId"`
	IconType     string `json:"iconType"`
	IconUrl      string `json:"iconUrl"`
	CoverUrl     string `json:"cover"`
	CoverType    string `json:"coverType"`
	Order        int    `json:"order"`
	ParentId     string `json:"parentId"`
	Title        string `json:"title"`
	Type         string `json:"type"`
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
func (p PageEntity) ChangePageEntityUrl(iconUrl string, coverUrl string) PageEntity {
	return PageEntity{
		Id:           p.Id,
		CurriculumId: p.CurriculumId,
		IconType:     p.IconType,
		IconUrl:      iconUrl,
		CoverUrl:     coverUrl,
		CoverType:    p.CoverType,
		Order:        p.Order,
		ParentId:     p.ParentId,
		Title:        p.Title,
	}
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
) (*PageEntity, error) {
	if Type != "curriculum" && Type != "info" && Type != "answer" {
		return nil, fmt.Errorf("unexpected type: %s", Type)
	}
	return &PageEntity{
		Id:           Id,
		CurriculumId: CurriculumId,
		IconType:     IconType,
		IconUrl:      IconUrl,
		CoverUrl:     CoverUrl,
		CoverType:    CoverType,
		Order:        Order,
		ParentId:     ParentId,
		Title:        Title,
		Type:         Type,
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
