package domain

import "fmt"

type NotionResourceType int

const (
	NotionResource NotionResourceType = iota
	DB
	DBQuery
	Block
	Children
	Page
	ChildDatabase
)

func (n NotionResourceType) GetRequestQuery(id string) (url string, method string, body map[string]any) {
	switch n {
	case DBQuery:
		url = fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", id)
		method = "POST"
		body = map[string]any{
			"filter": map[string]any{
				"property": "published",
				"checkbox": map[string]any{
					"equals": true,
				},
			},
		}
	case DB:
		url = fmt.Sprintf("https://api.notion.com/v1/databases/%s", id)
		method = "GET"

	case Page:
		url = fmt.Sprintf("https://api.notion.com/v1/pages/%s", id)
		method = "GET"

	case Block:
		url = fmt.Sprintf("https://api.notion.com/v1/blocks/%s", id)
		method = "GET"

	case Children:
		url = fmt.Sprintf("https://api.notion.com/v1/blocks/%s/children", id)
		method = "GET"
	case ChildDatabase:
		url = fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", id)
		method = "POST"
	}
	return url, method, body
}
