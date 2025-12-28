package domain

type PageProperty struct {
	Id         string         `json:"id"`
	Icon       *IconProperty  `json:"icon"`
	Cover      *CoverProperty `json:"cover"`
	Properties struct {
		Title TitleProperty `json:"title"`
	} `json:"properties"`
	Object         string `json:"object"`
	LastEditedTime string `json:"last_edited_time"`
	InTrash        bool   `json:"in_trash"`
}

type ParagraphProperty struct {
	RichText []RichTextProperty `json:"rich_text"`
	Color    string             `json:"color"`
}

type ToDoProperty struct {
	RichText []RichTextProperty `json:"rich_text"`
	Color    string             `json:"color"`
	Checked  bool               `json:"checked"`
}

type HeaderProperty struct {
	RichText     []RichTextProperty `json:"rich_text"`
	Color        string             `json:"color"`
	IsToggleable bool               `json:"is_toggleable"`
}
type CalloutProperty struct {
	RichText []RichTextProperty `json:"rich_text"`
	Color    string             `json:"color"`
	Icon     IconProperty       `json:"icon"`
}

type ImageProperty struct {
	Caption []RichTextProperty `json:"caption"`
	File    *struct {
		Url string `json:"url"`
	} `json:"file"`
	External *struct {
		Url string `json:"url"`
	} `json:"external"`
	Type string `json:"type"`
}

type EmbedProperty struct {
	Caption []RichTextProperty `json:"caption"`
	Url     string             `json:"url"`
}
type TableProperty struct {
	HasColumnHeader bool `json:"has_column_header"`
	HasRowHeader    bool `json:"has_row_header"`
	TableWidth      int  `json:"table_width"`
}
type TableRowProperty struct {
	Cells [][]RichTextProperty `json:"cells"`
}

type LinkToPageProperty struct {
	PageId string `json:"page_id"`
	Type   string `json:"type"`
}
type CodeProperty struct {
	Caption  []RichTextProperty `json:"caption"`
	Language string             `json:"language"`
	RichText []RichTextProperty `json:"rich_text"`
}
type SyncedProperty struct {
	SyncedFrom *struct {
		BlockId string `json:"block_id"`
		Type    string `json:"type"`
	} `json:"synced_from"`
}
