package model

// Page contains everything about a page. This will be stored.
type Page struct {
	Name     string         `json:"name"`
	Slug     string         `json:"slug"`
	Updated  string         `json:"updated"`
	Created  string         `json:"created"`
	HTML     string         `json:"html"`
	Md       string         `json:"md"`
	Parent   string         `json:"parent"`
	Children []PageFragment `json:"children"`
}

// func (s *Env) Save(p *Page) (*Page, error) {
// 	p, err := s.db.SavePage(p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p, nil
// }
