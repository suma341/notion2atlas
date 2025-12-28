package domain

import (
	"time"
)

type DBQueryEntity interface {
	GetId() string
	GetLastEditedTime() string
	GetTime() (*time.Time, error)
	GetUpdate() bool
	CompareQueryEntityTime(q DBQueryEntity) (bool, error)
	GetTitle() string
	ToPageEntity() (*PageEntity, error)
}

type NtBlock interface {
	GetId() string
	CompareQueryEntityTime(n NtBlock) (bool, error)
	GetTime() (*time.Time, error)
}
