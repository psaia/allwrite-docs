package model

import (
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

var items = Fragments{
	&PageFragment{
		Order: 3,
		Name:  "Images!",
		Type:  "file",
		Slug:  "images",
	},
	&PageFragment{
		Order: 0,
		Name:  "Homepage",
		Type:  "file",
		Slug:  "",
	},
	&PageFragment{
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
	&PageFragment{
		Order: 1,
		Name:  "Only one deep",
		Type:  "file",
		Slug:  "another-sub-directory",
	},
	&PageFragment{
		Order: 0,
		Name:  "This is a deep file",
		Type:  "file",
		Slug:  "another-sub-directory/a-deeper-directory",
	},
}

func TestMenuSorting(t *testing.T) {
	sorted := PageTree(items)

	if len(sorted) != 4 {
		t.Error("There should only be 4.")
	}
	if sorted[0].Name != "Homepage" {
		t.Error("Sorting is out of order: " + sorted[0].Name)
	}
	if len(sorted[0].Children) != 0 {
		t.Error("Index 0 should have zero children.")
	}

	if sorted[1].Name != "Only one deep" {
		t.Error("Sorting is out of order: " + sorted[1].Name)
	}
	if len(sorted[1].Children) != 1 {
		t.Error("Index 1 should have 1 child.")
	}

	if sorted[2].Name != "A Sub Directory" {
		t.Error("Sorting is out of order: " + sorted[2].Name)
	}
	if len(sorted[2].Children) != 1 {
		t.Error("Index 2 should have 2 children.")
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
}
