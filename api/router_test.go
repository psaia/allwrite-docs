package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/LevInteractive/allwrite-docs/model"
)

var items = []model.Page{
	model.Page{
		Order: 0,
		Slug:  "apples",
	},
	model.Page{
		Order: 1,
		Slug:  "apples/oranges",
	},
	model.Page{
		Order: 2,
		Slug:  "apples/pineapple",
	},
	model.Page{
		Order: 3,
		Slug:  "apples/strawberry",
	},
	model.Page{
		Order: 4,
		Slug:  "blueberry",
	},
	model.Page{
		Order: 5,
		Slug:  "greentea",
	},
	model.Page{
		Order: 1,
		Slug:  "greentea/fruity",
	},
	model.Page{
		Order: 0,
		Slug:  "greentea/fruity/leaves",
	},
	model.Page{
		Order: 0,
		Slug:  "",
	},
}

func TestMenuSorting(t *testing.T) {
	var testdata []*model.Page
	for _, item := range items {
		testdata = append(testdata, &item)
	}
	sorted := SortPages(testdata)
	fmt.Println(json.NewEncoder(w).Encode(sorted))

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
