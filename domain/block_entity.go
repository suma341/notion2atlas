package domain

type BlockEntity struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	ParentId     string `json:"parentId"`
	CurriculumId string `json:"curriculumId"`
	PageId       string `json:"pageId"`
	Data         any    `json:"data"`
	Order        int    `json:"order"`
}

func (b BlockEntity) GetId() string {
	return b.Id
}
