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
}
