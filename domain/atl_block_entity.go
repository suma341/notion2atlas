package domain

type AtlBlockEntity struct {
	Id           string             `json:"id"`
	Type         string             `json:"type"`
	ParentId     string             `json:"parentId"`
	CurriculumId string             `json:"curriculumId"`
	PageId       string             `json:"pageId"`
	Data         AtlBlockEntityData `json:"data"`
	Order        int                `json:"order"`
}

type AtlBlockEntityData struct {
	Type            string                    `json:"type"`
	Synced          *SyncedEntity             `json:"synced,omitempty"`
	Paragraph       *AtlParagraphEntity       `json:"paragraph,omitempty"`
	Todo            *AtlTodoEntity            `json:"todo,omitempty"`
	Header          *AtlHeaderEntity          `json:"header,omitempty"`
	Image           *AtlImageEntity           `json:"image,omitempty"`
	Embed           *AtlEmbedEntity           `json:"embed,omitempty"`
	Bookmark        *AtlBookmarkEntity        `json:"bookmark,omitempty"`
	Table           *TableEntity              `json:"table,omitempty"`
	TableRow        *AtlTableRowEntity        `json:"table_row,omitempty"`
	ChildPage       *ChildPageEntity          `json:"child_page,omitempty"`
	LinkToPage      *LinkToPageEntity         `json:"link_to_page,omitempty"`
	Code            *AtlCodeEntity            `json:"code,omitempty"`
	Callout         *AtlCalloutEntity         `json:"callout,omitempty"`
	ChildDB         *ChildDBEntity            `json:"child_database,omitempty"`
	TableOfContents *AtlTableOfContentsEntity `json:"table_of_contents,omitempty"`
}

type AtlTableOfContentsEntity []HeaderInfo

type HeaderInfo struct {
	HeaderType int    `json:"header_type"`
	BlockId    string `json:"block_id"`
	Text       string `json:"text"`
}

type AtlParagraphEntity struct {
	Color  string              `json:"color"`
	Parent []AtlRichTextEntity `json:"parent"`
}

type AtlTodoEntity struct {
	Color   string              `json:"color"`
	Parent  []AtlRichTextEntity `json:"parent"`
	Checked bool                `json:"checked"`
}

type AtlHeaderEntity struct {
	Color        string              `json:"color"`
	Parent       []AtlRichTextEntity `json:"parent"`
	IsToggleable bool                `json:"is_toggleable"`
}

func (h AtlHeaderEntity) GetComb() string {
	plain_text := ""
	for _, s := range h.Parent {
		plain_text = plain_text + s.PlainText
	}
	return plain_text
}

type AtlImageEntity struct {
	Parent []AtlRichTextEntity `json:"parent"`
	Url    string              `json:"url"`
	Width  int                 `json:"width"`
	Height int                 `json:"height"`
}

type AtlEmbedEntity struct {
	Parent   []AtlRichTextEntity `json:"parent"`
	Url      string              `json:"url"`
	CanEmbed bool                `json:"canEmbed"`
}

type AtlBookmarkEntity struct {
	Parent []AtlRichTextEntity `json:"parent"`
	Url    string              `json:"url"`
	Ogp    OGPResult           `json:"ogp"`
}

type AtlTableRowEntity [][]AtlRichTextEntity

type AtlCodeEntity struct {
	Language string              `json:"language"`
	Caption  []AtlRichTextEntity `json:"caption"`
	Parent   []AtlRichTextEntity `json:"parent"`
}

type AtlCalloutEntity struct {
	Color  string              `json:"color"`
	Parent []AtlRichTextEntity `json:"parent"`
	Icon   IconProperty        `json:"icon"`
}

type AtlRichTextEntity struct {
	Annotations AnnotationsProperty `json:"annotations"`
	PlainText   string              `json:"plain_text"`
	Href        *string             `json:"href"`
	Scroll      *string             `json:"scroll,omitempty"`
	Mention     *MentionEntity      `json:"mention,omitempty"`
	IsSameBP    bool                `json:"is_same_bp"`
}
