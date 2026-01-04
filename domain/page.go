package domain

import "time"

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
	CompareQueryEntityTime(n NtBlock) (bool, error)
	GetTime() (*time.Time, error)
}

type DBBasePage interface {
	GetCategories() []string
	GetVisilities() []string
	GetTags() []string
	GetId() string
}
