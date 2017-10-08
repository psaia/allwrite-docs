package model

import (
	"testing"
)

var items = Fragments{
	&PageFragment{ // top level
		Order: 0,
		Name:  "Homepage",
		Type:  "file",
		Slug:  "",
	},
	&PageFragment{ // top level
		Order: 3,
		Name:  "Images!",
		Type:  "file",
		Slug:  "images",
	},
	&PageFragment{ // top level
		Order: 2,
		Name:  "A Sub Directory",
		Type:  "dir",
		Slug:  "a-sub-directory",
	},
	&PageFragment{
		Order: 1,
		Name:  "How to be a friend",
		Type:  "file",
		Slug:  "a-sub-directory/how-to-be-a-friend",
	},
	&PageFragment{ // top level
		Order: 1,
		Name:  "I go deep",
		Type:  "file",
		Slug:  "deep",
	},
	&PageFragment{
		Order: 0,
		Name:  "This is a deep file",
		Type:  "file",
		Slug:  "deep/a-deeper-directory",
	},
	&PageFragment{
		Order: 0,
		Name:  "The deepest file",
		Type:  "file",
		Slug:  "deep/a-deeper-directory/hey-you",
	},
	&PageFragment{
		Order: 1,
		Name:  "A sibling to the deepest",
		Type:  "file",
		Slug:  "deep/a-deeper-directory/foobar",
	},
}

func TestMenuSorting(t *testing.T) {
	sorted := PageTree(items)

	// for _, page := range sorted {
	// 	fmt.Println(page.Name)
	// }
	if len(sorted) != 4 {
		t.Error("There should only be 4.")
	}
	if sorted[0].Name != "Homepage" {
		t.Error("Sorting is out of order: " + sorted[0].Name)
	}
	if len(sorted[0].Children) != 0 {
		t.Error("Index 0 should have zero children.")
	}

	if sorted[1].Name != "I go deep" {
		t.Error("Sorting is out of order: " + sorted[1].Name)
	}
	if len(sorted[1].Children) != 1 {
		t.Error("Index 1 should have 1 child.")
	}
	if len(sorted[1].Children[0].Children) != 2 {
		t.Error("This sub directory should have 2 deeper children.")
	}

	if sorted[2].Name != "A Sub Directory" {
		t.Error("Sorting is out of order: " + sorted[2].Name)
	}
	if len(sorted[2].Children) != 1 {
		t.Error("Index 2 should have 1 child.")
	}

	if sorted[3].Name != "Images!" {
		t.Error("Sorting is out of order: " + sorted[3].Name)
	}
	if len(sorted[3].Children) != 0 {
		t.Error("Index 3 should have no children")
	}
	if sorted[1].Children[0].Name != "This is a deep file" {
		t.Error("Index 1 has the wrong child.")
	}

	if sorted[3].Name != "Images!" {
		t.Error("Sorting is out of order: " + sorted[3].Name)
	}
}
