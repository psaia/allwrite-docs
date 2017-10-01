package model

// Page contains everything about a page. This will be stored.
type Page struct {
	PageFragment
	DocID    string          `json:"doc_id"`
	HTML     string          `json:"html"`
	Md       string          `json:"md"`
	Children []*PageFragment `json:"children,omitempty"`
}

// Pages is a slide of pages.
type Pages []*Page
