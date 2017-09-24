package model

// Page contains everything about a page. This will be stored.
type Page struct {
	Name     string         `json:"name"`
	DocID    string         `json:"doc_id"`
	Slug     string         `json:"slug"`
	Updated  string         `json:"updated"`
	Created  string         `json:"created"`
	Order    int            `json:"order"`
	HTML     string         `json:"html"`
	Md       string         `json:"md"`
	Parent   string         `json:"parent"`
	Children []PageFragment `json:"children"`
}
