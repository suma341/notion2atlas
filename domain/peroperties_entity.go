package domain

type RichTextProperty = struct {
	Annotations AnnotationsProperty `json:"annotations"`
	Href        *string             `json:"href"`
	PlainText   string              `json:"plain_text"`
	Mention     *MentionProperty    `json:"mention,omitempty"`
}

type AnnotationsProperty = struct {
	Bold          bool   `json:"bold"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
}

type MentionProperty = struct {
	LinkMention *LinkMentionProperty `json:"link_mention,omitempty"`
	Page        *PageProperty        `json:"page,omitempty"`
	Type        string               `json:"type"`
}

type LinkMentionProperty = struct {
	Href         string `json:"href"`
	IconUrl      string `json:"icon_url"`
	LinkProvider string `json:"link_provider"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Title        string `json:"title"`
}

type IconProperty struct {
	Type        string       `json:"type"`
	External    *UrlProperty `json:"external,omitempty"`
	File        *UrlProperty `json:"file,omitempty"`
	Emoji       *string      `json:"emoji,omitempty"`
	CustomEmoji *UrlProperty `json:"custom_emoji,omitempty"`
}

func (i IconProperty) GetIconUrl() string {
	iconType := i.Type
	iconUrl := ""
	switch iconType {
	case "external":
		if i.External != nil {
			iconUrl = i.External.Url
		}
	case "file":
		if i.File != nil {
			iconUrl = i.File.Url
		}
	case "custom_emoji":
		if i.CustomEmoji != nil {
			iconUrl = i.CustomEmoji.Url
		}
	case "emoji":
		if i.Emoji != nil {
			iconUrl = *i.Emoji
		}
	}
	return iconUrl
}

type CoverProperty struct {
	Type     string       `json:"type"`
	External *UrlProperty `json:"external,omitempty"`
	File     *UrlProperty `json:"file,omitempty"`
}

func (c CoverProperty) GetCoverUrl() string {
	coverUrl := ""
	coverType := c.Type
	switch coverType {
	case "external":
		if c.External != nil {
			coverUrl = c.External.Url
		}
	case "file":
		if c.File != nil {
			coverUrl = c.File.Url
		}
	}
	return coverUrl
}

type UrlProperty struct {
	Url string `json:"url"`
}
