package domain

type PageIf interface {
	GetId() string
	GetIcon() (IconType string, IconUrl string)
	GetCover() (CoverType string, CoverUrl string)
	GetTitle() string
}

type BasePage interface {
	GetId() string
	GetTitle() string
	ToPageEntity() (*PageEntity, error)
}
