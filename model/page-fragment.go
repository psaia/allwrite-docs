package model

// PageFragment is
type PageFragment struct {
	Name     string          `json:"name"`
	Slug     string          `json:"slug"`
	Updated  string          `json:"updated"`
	Created  string          `json:"created"`
	Children []*PageFragment `json:"children"`
}
