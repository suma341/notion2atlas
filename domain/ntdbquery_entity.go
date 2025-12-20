package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NtDBQueryEntity struct {
	Id             string         `json:"id"`
	Properties     map[string]any `json:"properties"`
	Icon           IconProperty   `json:"icon"`
	Cover          *CoverProperty `json:"cover"`
	LastEditedTime string         `json:"last_edited_time"`
	Object         string         `json:"object"` // list
}

func Res2NtDBQueryEntity(results []any) (*[]NtDBQueryEntity, error) {
	var filtered []NtDBQueryEntity
	for _, item := range results {
		objMap, ok := item.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("type convert fiailed: item.(map[string]any)")
		}
		jsonBytes, err := json.Marshal(objMap)
		if err != nil {
			fmt.Println("error in domain/Res2NtDBQueryEntity/json.Marshal(objMap)")
			return nil, err
		}
		var obj NtObjectProbe
		err = json.Unmarshal(jsonBytes, &obj)
		if err != nil {
			fmt.Println("error in domain/Res2NTBlockEntity/son.Unmarshal(jsonBytes, &obj)")
			return nil, err
		}
		if obj.Object != "error" {
			var block NtDBQueryEntity
			if err := json.Unmarshal(jsonBytes, &block); err != nil {
				return nil, err
			}
			block.Id = strings.ReplaceAll(block.Id, "-", "")
			filtered = append(filtered, block)
		} else {
			var ntErr NTErrorEntity
			if err := json.Unmarshal(jsonBytes, &ntErr); err != nil {
				return nil, err
			}
		}
	}
	return &filtered, nil
}

func (ntdbq NtDBQueryEntity) ToCurriculumEntity() (*CurriculumEntity, error) {
	id := ntdbq.Id
	pro, err := Map2Struct[CurriculumProperties](ntdbq.Properties)
	if err != nil {
		fmt.Println("error in converter/Query2Curriculum/Map2Struct")
		return nil, err
	}
	title := pro.Title.GetCombinedPlainText()
	var tag = []string{}
	if pro.Tag != nil {
		var tags = pro.Tag.MultiSelect
		for _, tag_item := range tags {
			tag = append(tag, tag_item.Name)
		}
	}
	var visibility = []string{}
	if pro.Visibility != nil {
		var visibilities = pro.Visibility.MultiSelect
		for _, visibility_item := range visibilities {
			visibility = append(visibility, visibility_item.Name)
		}
	}
	var order = pro.Order.Number
	var category = pro.Category.Select.Name
	var iconType = ntdbq.Icon.Type
	var iconUrl = ntdbq.Icon.GetIconUrl()
	var coverUrl = ""
	var coverType = ""
	if ntdbq.Cover != nil {
		coverType = ntdbq.Cover.Type
		coverUrl = ntdbq.Cover.GetCoverUrl()
	}
	var curriculum = CurriculumEntity{
		Id:             id,
		Title:          title,
		Tag:            tag,
		Visibility:     visibility,
		Order:          order,
		Category:       category,
		IconType:       iconType,
		IconUrl:        iconUrl,
		CoverType:      coverType,
		CoverUrl:       coverUrl,
		LastEditedTime: ntdbq.LastEditedTime,
		Update:         pro.Update.Checkbox,
	}
	return &curriculum, nil
}

func (ntdbq NtDBQueryEntity) ToInfoEntity() (*InfoEntity, error) {
	id := ntdbq.Id
	pro, err := Map2Struct[InfoProperties](ntdbq.Properties)
	if err != nil {
		fmt.Println("error in converter/Query2InfoEntity/Map2Struct")
		return nil, err
	}
	title := pro.Title.GetCombinedPlainText()
	var order = pro.Order.Number
	var iconType = ntdbq.Icon.Type
	var iconUrl = ntdbq.Icon.GetIconUrl()
	var coverUrl = ""
	var coverType = ""
	if ntdbq.Cover != nil {
		coverType = ntdbq.Cover.Type
		coverUrl = ntdbq.Cover.GetCoverUrl()
	}
	return &InfoEntity{
		Id:             id,
		Title:          title,
		Order:          order,
		IconType:       iconType,
		IconUrl:        iconUrl,
		CoverType:      coverType,
		CoverUrl:       coverUrl,
		LastEditedTime: ntdbq.LastEditedTime,
		Update:         pro.Update.Checkbox,
	}, nil
}

func (ntdbq NtDBQueryEntity) ToAnswerEntity() (*AnswerEntity, error) {
	id := ntdbq.Id
	pro, err := Map2Struct[AnswerProperties](ntdbq.Properties)
	if err != nil {
		fmt.Println("error in converter/Query2AnswerEntity/Map2Struct")
		return nil, err
	}
	title := pro.Title.GetCombinedPlainText()
	var order = pro.Order.Number
	var iconType = ntdbq.Icon.Type
	var iconUrl = ntdbq.Icon.GetIconUrl()
	var coverUrl = ""
	var coverType = ""
	if ntdbq.Cover != nil {
		coverType = ntdbq.Cover.Type
		coverUrl = ntdbq.Cover.GetCoverUrl()
	}
	return &AnswerEntity{
		Id:             id,
		Title:          title,
		Order:          order,
		IconType:       iconType,
		IconUrl:        iconUrl,
		CoverType:      coverType,
		CoverUrl:       coverUrl,
		LastEditedTime: ntdbq.LastEditedTime,
		Update:         pro.Update.Checkbox,
	}, nil
}

func (ntdbq NtDBQueryEntity) ToCategoryEntity() (*CategoryEntity, error) {
	id := ntdbq.Id
	pro, err := Map2Struct[CategoryProperties](ntdbq.Properties)
	if err != nil {
		fmt.Println("error in domain/ToCategoryEntity/Map2Struct")
		return nil, err
	}
	title := pro.Title.GetCombinedPlainText()
	description := pro.Description.GetCombinedPlainText()
	var order = pro.Order.Number
	var iconType = ntdbq.Icon.Type
	var iconUrl = ntdbq.Icon.GetIconUrl()
	var coverUrl = ""
	var coverType = ""
	if ntdbq.Cover != nil {
		coverType = ntdbq.Cover.Type
		coverUrl = ntdbq.Cover.GetCoverUrl()
	}
	var category = CategoryEntity{
		Id:                id,
		Title:             title,
		Description:       description,
		Order:             order,
		IconType:          iconType,
		IconUrl:           iconUrl,
		CoverType:         coverType,
		CoverUrl:          coverUrl,
		IsBasicCurriculum: pro.IsBasicCurriculum.Checkbox,
		LastEditedTime:    ntdbq.LastEditedTime,
		Update:            pro.Update.Checkbox,
	}
	return &category, nil
}
