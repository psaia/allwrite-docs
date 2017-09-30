package model

import (
	"regexp"
	"sort"
	"strings"
)

// PageFragment is a simple version of a page.
type PageFragment struct {
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Slug     string          `json:"slug"`
	Order    int             `json:"order"`
	Updated  string          `json:"updated"`
	Created  string          `json:"created"`
	Children []*PageFragment `json:"children"`
}

// Fragments is a slice of frags.
type Fragments []*PageFragment

// ByOrder will sort by order.
type ByOrder Fragments

func (slice ByOrder) Len() int {
	return len(slice)
}

func (slice ByOrder) Less(i, j int) bool {
	return slice[i].Order < slice[j].Order
}

func (slice ByOrder) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Only match the end of the slug: /foo/bar(/last)
var endPattern = regexp.MustCompile("\\/[^/]*$")

// Recursively fill children and order. This could be a little more elegant and
// performant I think.
func getChildrenForPageBySlug(curPage *PageFragment, pages Fragments) Fragments {
	slug := curPage.Slug
	children := make([]*PageFragment, 0, len(pages))
	for idx, page := range pages {
		b := endPattern.ReplaceAllString(page.Slug, "")
		if b == slug && page != curPage {
			pagesNext := make(Fragments, len(pages))
			copy(pagesNext, pages)
			sort.Sort(ByOrder(pagesNext))
			pagesNext = append(pagesNext[:idx], pagesNext[idx+1:]...)
			children = append(children, &PageFragment{
				Name:     page.Name,
				Slug:     page.Slug,
				Order:    page.Order,
				Type:     page.Type,
				Updated:  page.Updated,
				Created:  page.Created,
				Children: getChildrenForPageBySlug(curPage, pagesNext),
			})
		}
	}
	return children
}

// PageTree takes in a collection of pages and sorts them by order and populates the
// children slice.
func PageTree(pages Fragments) Fragments {
	sort.Sort(ByOrder(pages))
	master := make(Fragments, 0, len(pages))

	for _, page := range pages {
		segs := strings.Split(page.Slug, "/")
		depth := len(segs) - 1

		if depth == 0 {
			frag := &PageFragment{
				Name:     page.Name,
				Slug:     page.Slug,
				Order:    page.Order,
				Type:     page.Type,
				Updated:  page.Updated,
				Created:  page.Created,
				Children: getChildrenForPageBySlug(page, pages),
			}
			master = append(master, frag)
		}
	}
	return master
}
