package utils

import (
	"strings"
)

type UrlParam struct {
	Href   string
	Scroll *string
}

func RewriteHref(href string) UrlParam {
	if strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "http://") {
		if strings.HasPrefix(href, "https://www.notion.so") {
			pageHref := strings.Split(href, "/")[3]
			pageParams := strings.Split(pageHref, "#")
			pageId := pageParams[0]
			url := "/posts/curriculums/" + pageId
			if len(pageParams) > 1 {
				return UrlParam{
					Href:   url,
					Scroll: &pageParams[1],
				}
			}
			return UrlParam{Href: url}
		} else {
			return UrlParam{Href: href}
		}
	} else {
		pageParams := strings.Split(href, "#")
		pageId := pageParams[0][1:]
		url := "/posts/curriculums/" + pageId
		if len(pageParams) > 1 {
			return UrlParam{
				Href:   url,
				Scroll: &pageParams[1],
			}
		} else {
			return UrlParam{Href: url}
		}
	}
}
