package model

import (
	"strings"
)

// Page contains everything about a page. This will be stored.
type Page struct {
	Name     string         `json:"name"`
	Type     string         `json:"type"`
	DocID    string         `json:"doc_id"`
	Slug     string         `json:"slug"`
	Updated  string         `json:"updated"`
	Created  string         `json:"created"`
	Order    int            `json:"order"`
	HTML     string         `json:"html"`
	Md       string         `json:"md"`
	Children []PageFragment `json:"children"`
}

// Pages is a slide of pages.
type Pages []*Page

func (slice Pages) Len() int {
	return len(slice)
}

func (slice Pages) Less(i, j int) bool {
	return slice[i].Order < slice[j].Order
}

func (slice Pages) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// PageTree takes in a collection of pages and sorts them by order and populates the
// children slice.
func PageTree(pages Pages) Pages {
	var maxDepth, n, currentDepth int

	for idx := range pages {
		l := len(strings.Split(pages[idx].Slug, "/"))
		if l > n {
			maxDepth = l
		} else {
			n = l
		}
	}

	master := make(Pages, 0, len(pages))

	for currentDepth < maxDepth {
		for _, page := range pages {
			segs := strings.Split(page.Slug, "/")
			l := len(segs)
			if currentDepth == 0 && l == 1 {
				master = append(master, page)
			} else if currentDepth > 0 && currentDepth == l-1 {
				// wat
			}
		}
		currentDepth++
	}

	return master
}
