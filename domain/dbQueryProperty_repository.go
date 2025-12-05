package domain

type TitleProperty struct {
	Title []RichTextProperty `json:"title"`
}

type TextProperty struct {
	RichText []RichTextProperty `json:"rich_text"`
}

func (t TitleProperty) GetCombinedPlainText() string {
	plain_text := ""
	for _, s := range t.Title {
		plain_text = plain_text + s.PlainText
	}
	return plain_text
}

func (t TextProperty) GetCombinedPlainText() string {
	plain_text := ""
	for _, s := range t.RichText {
		plain_text = plain_text + s.PlainText
	}
	return plain_text
}

type ChackBoxQuery struct {
	Checkbox bool `json:"checkbox"`
}

type MultiSelectQuery struct {
	MultiSelect []struct {
		Name string `json:"name"`
	} `json:"multi_select"`
}
type NumberQuery struct {
	Number int
}
type SelectQuery struct {
	Select struct {
		Name string `json:"name"`
	} `json:"select"`
}
