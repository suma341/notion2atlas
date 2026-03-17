package domain

type BlockEntity struct {
	Id           string          `json:"id"`
	Type         string          `json:"type"`
	ParentId     string          `json:"parentId"`
	CurriculumId string          `json:"curriculumId"`
	PageId       string          `json:"pageId"`
	Data         BlockEntityData `json:"data"`
	Order        int             `json:"order"`
}

func (b BlockEntity) GetId() string {
	return b.Id
}

func (b BlockEntity) ToAtlEntity(data AtlBlockEntityData) AtlBlockEntity {
	return AtlBlockEntity{
		Id:           b.Id,
		Type:         b.Type,
		ParentId:     b.ParentId,
		CurriculumId: b.CurriculumId,
		PageId:       b.PageId,
		Data:         data,
		Order:        b.Order,
	}
}

type BlockEntityData struct {
	Type       string            `json:"type"`
	Synced     *SyncedEntity     `json:"synced,omitempty"`
	Paragraph  *ParagraphEntity  `json:"paragraph,omitempty"`
	Todo       *TodoEntity       `json:"todo,omitempty"`
	Header     *HeaderEntity     `json:"header,omitempty"`
	Image      *ImageEntity      `json:"image,omitempty"`
	Embed      *EmbedEntity      `json:"embed,omitempty"`
	Bookmark   *BookmarkEntity   `json:"bookmark,omitempty"`
	Table      *TableEntity      `json:"table,omitempty"`
	TableRow   *TableRowEntity   `json:"table_row,omitempty"`
	ChildPage  *ChildPageEntity  `json:"child_page,omitempty"`
	LinkToPage *LinkToPageEntity `json:"link_to_page,omitempty"`
	Code       *CodeEntity       `json:"code,omitempty"`
	Callout    *CalloutEntity    `json:"callout,omitempty"`
	ChildDB    *ChildDBEntity    `json:"child_database,omitempty"`
}

func (b BlockEntityData) ToAtlData(parent *[]AtlRichTextEntity) AtlBlockEntityData {
	switch b.Type {
	case "paragraph":
		atlEntity := b.Paragraph.ToAtl(parent)
		return AtlBlockEntityData{
			Type:      b.Type,
			Paragraph: &atlEntity,
		}
	case "todo":
		atlEntity := b.Todo.ToAtl(parent)
		return AtlBlockEntityData{
			Type: b.Type,
			Todo: &atlEntity,
		}
	case "header":
		atlEntity := b.Header.ToAtl(parent)
		return AtlBlockEntityData{
			Type:   b.Type,
			Header: &atlEntity,
		}
	case "image":
		atlEntity := b.Image.ToAtl(parent)
		return AtlBlockEntityData{
			Type:  b.Type,
			Image: &atlEntity,
		}
	case "embed":
		atlEntity := b.Embed.ToAtl(parent)
		return AtlBlockEntityData{
			Type:  b.Type,
			Embed: &atlEntity,
		}
	case "bookmark":
		atlEntity := b.Bookmark.ToAtl(parent)
		return AtlBlockEntityData{
			Type:     b.Type,
			Bookmark: &atlEntity,
		}
	case "table":
		return AtlBlockEntityData{
			Type:  b.Type,
			Table: b.Table,
		}
	case "table_row":
		atlEntity := b.TableRow.ToAtl(nil)
		return AtlBlockEntityData{
			Type:     b.Type,
			TableRow: &atlEntity,
		}
	case "child_page":
		return AtlBlockEntityData{
			Type:      b.Type,
			ChildPage: b.ChildPage,
		}
	case "link_to_page":
		return AtlBlockEntityData{
			Type:       b.Type,
			LinkToPage: b.LinkToPage,
		}
	case "callout":
		atlEntity := b.Callout.ToAtl(parent)
		return AtlBlockEntityData{
			Type:    b.Type,
			Callout: &atlEntity,
		}
	case "code":
		atlEntity := b.Code.ToAtl(nil, nil)
		return AtlBlockEntityData{
			Type: b.Type,
			Code: &atlEntity,
		}
	case "child_database":
		return AtlBlockEntityData{
			Type:    b.Type,
			ChildDB: b.ChildDB,
		}
	case "synced":
		return AtlBlockEntityData{
			Type:   b.Type,
			Synced: b.Synced,
		}
	default:
		return AtlBlockEntityData{
			Type: b.Type,
		}
	}
}

func (b BlockEntityData) GetHasParentEntity() HasParentBlock {
	switch b.Type {
	case "paragraph":
		return b.Paragraph
	case "todo":
		return b.Todo
	case "header":
		return b.Header
	case "image":
		return b.Image
	case "embed":
		return b.Embed
	case "bookmark":
		return b.Bookmark
	case "callout":
		return b.Callout
	default:
		return nil
	}
}

type IBlockEntity interface {
	GetType() string
}

type HasParentBlock interface {
	GetParent() []RichTextEntity
}

type ParagraphEntity struct {
	Color  string           `json:"color"`
	Parent []RichTextEntity `json:"parent"`
}

func (p ParagraphEntity) GetConcatenatedText() string {
	return concatenateRichTextsText(p.Parent)
}

func (p ParagraphEntity) GetType() string {
	return "paragraph"
}

func (p ParagraphEntity) GetParent() []RichTextEntity {
	return p.Parent
}

func (p ParagraphEntity) ToAtl(atlParents *[]AtlRichTextEntity) AtlParagraphEntity {
	atl := EntityArrToAtlArr(p.Parent)
	if atlParents != nil {
		atl = *atlParents
	}
	return AtlParagraphEntity{
		Color:  p.Color,
		Parent: atl,
	}
}

type TodoEntity struct {
	Color   string           `json:"color"`
	Parent  []RichTextEntity `json:"parent"`
	Checked bool             `json:"checked"`
}

func (p TodoEntity) GetConcatenatedText() string {
	return concatenateRichTextsText(p.Parent)
}

func (p TodoEntity) GetType() string {
	return "todo"
}

func (p TodoEntity) GetParent() []RichTextEntity {
	return p.Parent
}
func (p TodoEntity) ToAtl(atlParents *[]AtlRichTextEntity) AtlTodoEntity {
	atl := EntityArrToAtlArr(p.Parent)
	if atlParents != nil {
		atl = *atlParents
	}
	return AtlTodoEntity{
		Color:  p.Color,
		Parent: atl,
	}
}

type HeaderEntity struct {
	Color        string           `json:"color"`
	Parent       []RichTextEntity `json:"parent"`
	IsToggleable bool             `json:"is_toggleable"`
}

func (t HeaderEntity) GetCombinedPlainText() string {
	plain_text := ""
	for _, s := range t.Parent {
		plain_text = plain_text + s.PlainText
	}
	return plain_text
}

func (p HeaderEntity) GetConcatenatedText() string {
	return concatenateRichTextsText(p.Parent)
}

func (p HeaderEntity) GetType() string {
	return "header"
}
func (p HeaderEntity) GetParent() []RichTextEntity {
	return p.Parent
}
func (p HeaderEntity) ToAtl(atlParents *[]AtlRichTextEntity) AtlHeaderEntity {
	atl := EntityArrToAtlArr(p.Parent)
	if atlParents != nil {
		atl = *atlParents
	}
	return AtlHeaderEntity{
		Color:        p.Color,
		Parent:       atl,
		IsToggleable: p.IsToggleable,
	}
}

type ImageEntity struct {
	Parent []RichTextEntity `json:"parent"`
	Url    string           `json:"url"`
}

func (p ImageEntity) GetType() string {
	return "image"
}
func (p ImageEntity) GetParent() []RichTextEntity {
	return p.Parent
}
func (p ImageEntity) ToAtl(atlParents *[]AtlRichTextEntity) AtlImageEntity {
	atl := EntityArrToAtlArr(p.Parent)
	if atlParents != nil {
		atl = *atlParents
	}
	return AtlImageEntity{
		Url:    p.Url,
		Parent: atl,
	}
}

type EmbedEntity struct {
	Parent   []RichTextEntity `json:"parent"`
	Url      string           `json:"url"`
	CanEmbed bool             `json:"canEmbed"`
}

func (p EmbedEntity) GetType() string {
	return "embed"
}
func (p EmbedEntity) GetParent() []RichTextEntity {
	return p.Parent
}
func (p EmbedEntity) ToAtl(atlParents *[]AtlRichTextEntity) AtlEmbedEntity {
	atl := EntityArrToAtlArr(p.Parent)
	if atlParents != nil {
		atl = EntityArrToAtlArr(p.Parent)
	}
	return AtlEmbedEntity{
		Url:      p.Url,
		Parent:   atl,
		CanEmbed: p.CanEmbed,
	}
}

type BookmarkEntity struct {
	Parent []RichTextEntity `json:"parent"`
	Url    string           `json:"url"`
	Ogp    OGPResult        `json:"ogp"`
}

func (p BookmarkEntity) GetParent() []RichTextEntity {
	return p.Parent
}
func (p BookmarkEntity) GetType() string {
	return "bookmark"
}
func (p BookmarkEntity) ToAtl(atlParents *[]AtlRichTextEntity) AtlBookmarkEntity {
	atl := EntityArrToAtlArr(p.Parent)
	if atlParents != nil {
		atl = *atlParents
	}
	return AtlBookmarkEntity{
		Url:    p.Url,
		Parent: atl,
		Ogp:    p.Ogp,
	}
}

type TableEntity struct {
	HasColumnHeader bool `json:"has_column_header"`
	HasRowHeader    bool `json:"has_row_header"`
	TableWidth      int  `json:"table_width"`
}

func (p TableEntity) GetType() string {
	return "table"
}

type TableRowEntity [][]RichTextEntity

func (p TableRowEntity) ToAtl(atlParents *[][]AtlRichTextEntity) AtlTableRowEntity {
	atl := [][]AtlRichTextEntity{}
	if atlParents != nil {
		return *atlParents
	}
	for _, i := range p {
		atl = append(atl, EntityArrToAtlArr(i))
	}
	return atl
}

func (p TableRowEntity) GetType() string {
	return "table_row"
}

type ChildPageEntity struct {
	Parent   string `json:"parent"`
	IconType string `json:"iconType"`
	IconUrl  string `json:"iconUrl"`
	CoverUrl string `json:"coverUrl"`
}

func (p ChildPageEntity) GetType() string {
	return "child_page"
}

type LinkToPageEntity struct {
	Link     string `json:"link"`
	IconType string `json:"iconType"`
	IconUrl  string `json:"iconUrl"`
	Title    string `json:"title"`
}

func (p LinkToPageEntity) GetType() string {
	return "link_to_page"
}

type CodeEntity struct {
	Language string           `json:"language"`
	Caption  []RichTextEntity `json:"caption"`
	Parent   []RichTextEntity `json:"parent"`
}

func (p CodeEntity) ToAtl(cAtlParents *[]AtlRichTextEntity, pAtlParents *[]AtlRichTextEntity) AtlCodeEntity {
	c := EntityArrToAtlArr(p.Caption)
	if cAtlParents != nil {
		c = *cAtlParents
	}
	pa := EntityArrToAtlArr(p.Parent)
	if pAtlParents != nil {
		pa = *pAtlParents
	}
	return AtlCodeEntity{
		Language: p.Language,
		Caption:  c,
		Parent:   pa,
	}
}

func (p CodeEntity) GetType() string {
	return "code"
}
func (p CodeEntity) GetParent() []RichTextEntity {
	return p.Parent
}

type CalloutEntity struct {
	Color  string           `json:"color"`
	Parent []RichTextEntity `json:"parent"`
	Icon   IconProperty     `json:"icon"`
}

func (p CalloutEntity) GetConcatenatedText() string {
	return concatenateRichTextsText(p.Parent)
}

func (p CalloutEntity) GetType() string {
	return "callout"
}
func (p CalloutEntity) GetParent() []RichTextEntity {
	return p.Parent
}
func (p CalloutEntity) ToAtl(atlParents *[]AtlRichTextEntity) AtlCalloutEntity {
	pa := EntityArrToAtlArr(p.Parent)
	if atlParents != nil {
		pa = *atlParents
	}
	return AtlCalloutEntity{
		Color:  p.Color,
		Icon:   p.Icon,
		Parent: pa,
	}
}

type SyncedEntity string

func (p SyncedEntity) GetType() string {
	return "synced"
}

type ChildDBEntity struct {
	DatabaseData *NtDBEntity       `json:"database_data"`
	QueryData    []NtDBQueryEntity `json:"query_data"`
}

func (p ChildDBEntity) GetType() string {
	return "child_database"
}

type RichTextEntity struct {
	Annotations AnnotationsProperty `json:"annotations"`
	PlainText   string              `json:"plain_text"`
	Href        *string             `json:"href"`
	Scroll      *string             `json:"scroll,omitempty"`
	Mention     *MentionEntity      `json:"mention,omitempty"`
}

func concatenateRichTextsText(richTexts []RichTextEntity) string {
	concatenatedText := ""
	for _, r := range richTexts {
		concatenatedText = concatenatedText + r.PlainText
	}
	return concatenatedText
}

func (r RichTextEntity) ToAtlEntity(IsSameBP bool) AtlRichTextEntity {
	return AtlRichTextEntity{
		Annotations: r.Annotations,
		PlainText:   r.PlainText,
		Href:        r.Href,
		Scroll:      r.Scroll,
		Mention:     r.Mention,
		IsSameBP:    IsSameBP,
	}
}

func EntityArrToAtlArr(r []RichTextEntity) []AtlRichTextEntity {
	atl := []AtlRichTextEntity{}
	for _, i := range r {
		atl = append(atl, i.ToAtlEntity(false))
	}
	return atl
}

type MentionEntity = struct {
	Content map[string]any `json:"content"`
	Type    string         `json:"type"`
}

type OGPResult struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Icon        string `json:"icon"`
}
