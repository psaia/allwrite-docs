package model

import (
	"regexp"
	"sort"
	"strings"
)

// Page contains everything about a page. This will be stored.
type Page struct {
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	DocID    string          `json:"doc_id"`
	Slug     string          `json:"slug"`
	Updated  string          `json:"updated"`
	Created  string          `json:"created"`
	Order    int             `json:"order"`
	HTML     string          `json:"html"`
	Md       string          `json:"md"`
	Children []*PageFragment `json:"children"`
}

// Pages is a slide of pages.
type Pages []*Page

// BySlashLen is for sorting by the depth.
type BySlashLen []*Page

// ByOrder will sort by order.
type ByOrder []*Page

func (slice ByOrder) Len() int {
	return len(slice)
}

func (slice ByOrder) Less(i, j int) bool {
	return slice[i].Order < slice[j].Order
}

func (slice ByOrder) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice BySlashLen) Len() int {
	return len(slice)
}

func (slice BySlashLen) Less(i, j int) bool {
	ic := len(strings.Split(slice[i].Slug, "/"))
	jc := len(strings.Split(slice[j].Slug, "/"))
	return ic < jc
}

func (slice BySlashLen) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Only match the end of the slug: /foo/bar(/last)
var endPattern = regexp.MustCompile("\\/[^/]*$")

// Recursively fill children and order.
func getChildrenForPageBySlug(curPage *Page, pages Pages) []*PageFragment {
	slug := curPage.Slug
	children := make([]*PageFragment, 0, len(pages))
	for idx, page := range pages {
		b := endPattern.ReplaceAllString(page.Slug, "")
		if b == slug && page != curPage {
			pagesNext := make(Pages, len(pages))
			copy(pagesNext, pages)
			sort.Sort(ByOrder(pagesNext))
			pagesNext = append(pagesNext[:idx], pagesNext[idx+1:]...)
			children = append(children, &PageFragment{
				Name:     page.Name,
				Slug:     page.Slug,
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
func PageTree(pages Pages) Pages {
	sort.Sort(ByOrder(pages))
	master := make(Pages, 0, len(pages))

	for _, page := range pages {
		segs := strings.Split(page.Slug, "/")
		depth := len(segs) - 1

		if depth == 0 {
			master = append(master, page)
			page.Children = getChildrenForPageBySlug(page, pages)
		}
	}
	return master
}
