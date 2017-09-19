package model

// PageFragment is
type PageFragment struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Slug     string         `json:"slug"`
	Updated  uint64         `json:"updated"`
	Created  uint64         `json:"created"`
	Children []PageFragment `json:"children"`
}
