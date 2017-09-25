package model

import (
	"fmt"
	"testing"
)

// Some real sample data:
//
//         title        | type |                   slug                   | placement
// ---------------------+------+------------------------------------------+-----------
//  Images!             | file | images                                   |         2
//  Homepage            | file |                                          |         0
//  A Sub Directory     | dir  | a-sub-directory                          |         1
//  How to be a friend  | file | a-sub-directory/how-to-be-a-friend       |         1
//  Only one deep       | file | another-sub-directory                    |         0
//  This is a deep file | file | another-sub-directory/a-deeper-directory |         0

var items = Pages{
	&Page{
		Order: 2,
		Name:  "Images!",
		Type:  "file",
		Slug:  "images",
	},
	&Page{
		Order: 0,
		Name:  "Homepage",
		Type:  "file",
		Slug:  "",
	},
	&Page{
		Order: 1,
		Name:  "A Sub Directory",
		Type:  "dir",
		Slug:  "a-sub-directory",
	},
	&Page{
		Order: 1,
		Name:  "How to be a friend",
		Type:  "file",
		Slug:  "a-sub-directory/how-to-be-a-friend",
	},
	&Page{
		Order: 0,
		Name:  "Only one deep",
		Type:  "file",
		Slug:  "another-sub-directory",
	},
	&Page{
		Order: 0,
		Name:  "This is a deep file",
		Type:  "file",
		Slug:  "another-sub-directory/a-deeper-directory",
	},
}

func TestMenuSorting(t *testing.T) {
	sorted := PageTree(items)
	// sorted := items

	// json, err := json.Marshal(sorted)
	// if err != nil {
	// 	t.Error(err)
	// }
	for _, s := range sorted {
		fmt.Println(s)
	}

	// mdDoc := getFixture("./fixtures/sample-doc.md")
	// htmlDoc := getFixture("./fixtures/sample-doc.html")
	//
	// r := strings.NewReader(htmlDoc)
	// transformedMd, err := MarshalMarkdownFromHTML(r)
	//
	// if err != nil {
	// 	t.Error(err)
	// }
	//
	// if transformedMd != mdDoc {
	// 	t.Error("HTML did not translate to markdown properly.")
	// }
}
