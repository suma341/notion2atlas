package usecase

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/utils"
)

type RichTextModel struct {
	Annotations domain.AnnotationsProperty `json:"annotations"`
	PlainText   string                     `json:"plain_text"`
	Href        *string                    `json:"href"`
	Scroll      *string                    `json:"scroll,omitempty"`
	Mention     *MentionModel              `json:"mention,omitempty"`
}

type MentionModel = struct {
	Content map[string]any `json:"content"`
	Type    string         `json:"type"`
}

func ProcessRichText(richTexts []domain.RichTextProperty) ([]RichTextModel, error) {
	richTextModels := make([]RichTextModel, 0, len(richTexts))
	for _, it := range richTexts {
		model, err := richTextRes2Model(it)
		if err != nil {
			fmt.Println("error in usecase/ProcessRichText/richTextRes2Model")
			return nil, err
		}
		richTextModels = append(richTextModels, *model)
	}
	return richTextModels, nil
}

func richTextRes2Model(rich_text domain.RichTextProperty) (*RichTextModel, error) {
	var href *string = nil
	var scroll *string = nil
	if rich_text.Href != nil {
		urlParam := utils.RewriteHref(*rich_text.Href)
		href = &urlParam.Href
		scroll = urlParam.Scroll
	}
	if rich_text.Mention != nil {
		mentionModel, err := getMentionModel(rich_text.Mention)
		if err != nil {
			fmt.Println("error in usecase/richTextRes2Model/getMentionModel")
			return nil, err
		}
		return &RichTextModel{
			Annotations: rich_text.Annotations,
			PlainText:   rich_text.PlainText,
			Href:        href,
			Scroll:      scroll,
			Mention:     mentionModel,
		}, nil
	}
	return &RichTextModel{
		Annotations: rich_text.Annotations,
		PlainText:   rich_text.PlainText,
		Href:        href,
		Scroll:      scroll,
		Mention:     nil,
	}, nil
}

func getMentionModel(mention *domain.MentionProperty) (*MentionModel, error) {
	mentionType := mention.Type
	switch mentionType {
	case "page":
		if mention.Page != nil {
			pageId := mention.Page.Id
			pageData, err := GetPageItem(pageId)
			if err != nil {
				fmt.Println("error in usecase/getMentionModel/GetPageItem")
				return nil, err
			}
			return &MentionModel{
				Content: map[string]any{
					"iconUrl":  pageData.IconUrl,
					"iconType": pageData.IconType,
					"title":    pageData.Title,
				},
				Type: "prossedPage",
			}, nil
		}
	case "link_mention":
		var mentionContent, err = domain.Struct2Map(mention.LinkMention)
		if err != nil {
			fmt.Println("error in usecase/getMentionModel/domain.Struct2Map")
			return nil, err
		}
		return &MentionModel{
			Content: mentionContent,
			Type:    mentionType,
		}, nil
	}
	return nil, nil
}
