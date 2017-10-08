package model

import (
	"sort"
	"strings"
)

// Fragments is a slice of frags.
type Fragments []*PageFragment

// PageFragment is a simple version of a page.
type PageFragment struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Slug    string `json:"slug"`
	Order   int    `json:"order"`
	Updated string `json:"updated"`
	Created string `json:"created"`

	// Only populated on search results.
	MatchingText string `json:"reltext,omitempty"`

	// Children on only populated when querying the menu.
	Children Fragments `json:"children,omitempty"`
}

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

func compareSlices(a []string, b []string) int {
	n, aLen, bLen := 0, len(a), len(b)
	for i := range a {
		if i < aLen && i < bLen && a[i] == b[i] {
			n++
		}
	}
	return n
}

func appendDeep(
	master Fragments,
	prevSibling *Fragments,
	prev *PageFragment,
	next *PageFragment,
) (
	Fragments,
	*Fragments,
	*PageFragment,
) {
	prevSegs := strings.Split(prev.Slug, "/")
	nextSegs := strings.Split(next.Slug, "/")
	diff := compareSlices(prevSegs, nextSegs)
	var siblings *Fragments

	// These comments are super helpful when debugging. Leaving commented out.
	if len(prevSegs) == diff {
		// log.Printf("Adding child '%s' to parent '%s' -- diff: %v", next.Name, prev.Name, diff)
		prev.Children = append(prev.Children, next)
		siblings = &prev.Children
	} else if diff > 0 {
		// log.Printf("Adding sibling: '%s' along side '%s'", next.Name, prev.Name)
		*prevSibling = append(*prevSibling, next)
		siblings = prevSibling
	} else {
		// log.Printf("Adding top level: %s -- diff: %v", next.Name, diff)
		master = append(master, next)
		siblings = &master
	}

	sort.Sort(ByOrder(master))
	sort.Sort(ByOrder(*siblings))

	return master, siblings, next
}

// PageTree takes in an ordered collection of pages and sorts them by order and
// populates the children slice. These must be ordered by slug.
//
// For example:
//
//         title         |                     slug
// ----------------------+----------------------------------------------
//  Welcome              |
//  First Page           | getting-started
//  Code Snippets        | getting-started/code-snippets
//  Hello World          | getting-started/hello-world
//  Section One          | section-one
//  Moderately Deep File | section-one/moderately-deep-file
//  Only one deep        | section-two
//  This is a deep file  | section-two/subsection-one
//  Subsection Two       | section-two/subsection-two
//  Deep Page Example    | section-two/subsection-two/deep-page-example
func PageTree(pages Fragments) Fragments {
	master := make(Fragments, 0, len(pages))

	// Append the first page to get our master started. Then shift off the first
	// item of the main pages array. Singlings gets passed by pointer so siblings
	// can be added to it if need be.
	master = append(master, pages[0])
	siblings := &master
	pages = pages[1:]

	// This is set on each iteration.
	prev := master[0]

	// Loop through the rest of the pages adding them to the master where they
	// belong.
	for _, page := range pages {
		master, siblings, prev = appendDeep(master, siblings, prev, page)
	}

	return master
}
