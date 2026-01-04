package usecase

import (
	"errors"
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase/notionUC"
	"notion2atlas/utils"
)

func ProcessRichText(richTexts []domain.RichTextProperty, type_ string) ([]domain.RichTextEntity, error) {
	richTextModels := make([]domain.RichTextEntity, 0, len(richTexts))
	for _, it := range richTexts {
		model, err := richTextRes2Model(it, type_)
		if err != nil {
			if !errors.Is(err, domain.ErrNotionErrorResponse) {
				fmt.Println("error in usecase/ProcessRichText/richTextRes2Model")
				return nil, err
			}
		}
		richTextModels = append(richTextModels, *model)
	}
	return richTextModels, nil
}

func richTextRes2Model(rich_text domain.RichTextProperty, type_ string) (*domain.RichTextEntity, error) {
	var href *string = nil
	var scroll *string = nil
	if rich_text.Href != nil {
		urlParam := utils.RewriteHref(*rich_text.Href)
		href = &urlParam.Href
		scroll = urlParam.Scroll
	}
	if rich_text.Mention != nil {
		mentionModel, err := getMentionModel(rich_text.Mention, type_)
		if err != nil {
			if errors.Is(err, domain.ErrNotionErrorResponse) {
				return nil, domain.ErrNotionErrorResponse
			}
			fmt.Println("error in usecase/richTextRes2Model/getMentionModel")
			return nil, err
		}
		return &domain.RichTextEntity{
			Annotations: rich_text.Annotations,
			PlainText:   rich_text.PlainText,
			Href:        href,
			Scroll:      scroll,
			Mention:     mentionModel,
		}, nil
	}
	return &domain.RichTextEntity{
		Annotations: rich_text.Annotations,
		PlainText:   rich_text.PlainText,
		Href:        href,
		Scroll:      scroll,
		Mention:     nil,
	}, nil
}

func getMentionModel(mention *domain.MentionProperty, type_ string) (*domain.MentionEntity, error) {
	mentionType := mention.Type
	switch mentionType {
	case "page":
		if mention.Page != nil {
			pageId := mention.Page.Id
			pageData, err := notionUC.GetPageItem(pageId, type_)
			if err != nil {
				if errors.Is(err, domain.ErrNotionErrorResponse) {
					return nil, domain.ErrNotionErrorResponse
				}
				fmt.Println("error in usecase/getMentionModel/GetPageItem")
				return nil, err
			}
			return &domain.MentionEntity{
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
		return &domain.MentionEntity{
			Content: mentionContent,
			Type:    mentionType,
		}, nil
	}
	return nil, nil
}
