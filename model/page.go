package model

// Page contains everything about a page. This will be stored.
type Page struct {
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Updated string `json:"updated"`
	Created string `json:"created"`
	HTML    string `json:"html"`
	Md      string `json:"md"`
	Parent  string `json:"parent"`
}

// func (*Page) Save
