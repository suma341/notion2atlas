package domain

import (
	"fmt"
	"notion2atlas/constants"
)

type DatRType int

const (
	PAGE_DAT DatRType = iota
	SYNCED_DAT
	CURRICULUM_DAT
	CATEGORY_DAT
)

func (r DatRType) GetPath() (dat string, tmp string) {
	switch r {
	case PAGE_DAT:
		return constants.PAGE_DAT_PATH, constants.TMP_ALL_PAGE_PATH
	case SYNCED_DAT:
		return constants.SYNCED_DAT_PATH, constants.TMP_ALL_PAGE_PATH
	case CURRICULUM_DAT:
		return constants.CURRICULUM_DAT_PATH, constants.TMP_ALL_CURRICULUM_PATH
	case CATEGORY_DAT:
		return constants.CATEGORY_DAT_PATH, constants.TMP_ALL_CATEGORY_PATH
	}
	return "", ""
}

type ResourceType int

const (
	CURRICULUM ResourceType = iota
	PAGE
	CATEGORY
	INFO
	ANSWER
	SYNCED
	TMP_PAGE
	TMP_ALL_PAGE
	TMP_ALL_CATEGORY
)

func (r ResourceType) GetStr() string {
	switch r {
	case CURRICULUM:
		return "curriculum"
	case PAGE:
		return "page"
	case CATEGORY:
		return "category"
	case INFO:
		return "info"
	case ANSWER:
		return "answer"
	case SYNCED:
		return "synced"
	default:
		return ""
	}
}

func (r ResourceType) GetFilePathFromResourceType() (string, error) {
	switch r {
	case CURRICULUM:
		return constants.TMP_ALL_CURRICULUM_PATH, nil
	case PAGE:
		return constants.TMP_ALL_PAGE_PATH, nil
	case CATEGORY:
		return constants.TMP_ALL_CATEGORY_PATH, nil
	case INFO:
		return constants.INFO_PATH, nil
	case ANSWER:
		return constants.ANSWER_PATH, nil
	case SYNCED:
		return constants.SYNCED_PATH, nil
	case TMP_PAGE:
		return constants.TMP_PAGE_PATH, nil
	case TMP_ALL_PAGE:
		return constants.TMP_ALL_PAGE_PATH, nil
	case TMP_ALL_CATEGORY:
		return constants.TMP_ALL_CATEGORY_PATH, nil
	default:
		return "", fmt.Errorf("unexpected resourceType")
	}
}
