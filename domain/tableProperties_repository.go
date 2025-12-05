package domain

type CurriculumProperties struct {
	Title      TitleProperty     `json:"title"`
	Tag        *MultiSelectQuery `json:"tag"`
	Visibility *MultiSelectQuery `json:"visibility"`
	Order      NumberQuery       `json:"order"`
	Category   SelectQuery       `json:"category"`
	Update     ChackBoxQuery     `json:"update"`
}

type CategoryProperties struct {
	Title             TitleProperty `json:"title"`
	Description       TextProperty  `json:"description"`
	IsBasicCurriculum ChackBoxQuery `json:"is_basic_curriculum"`
	Order             NumberQuery   `json:"order"`
	Update            ChackBoxQuery `json:"update"`
}

type InfoProperties struct {
	Title  TitleProperty `json:"title"`
	Order  NumberQuery   `json:"order"`
	Update ChackBoxQuery `json:"update"`
}

type AnswerProperties struct {
	Title  TitleProperty `json:"title"`
	Order  NumberQuery   `json:"order"`
	Update ChackBoxQuery `json:"update"`
}
