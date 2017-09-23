package store

import "github.com/LevInteractive/allwrite-docs/model"

// Store interface for any storage client to implement.
// Eventually, we'll probably want to make search return a pagination struct
// with a cursor implementation. Search is also very simple. Could add a few
// more parameters eventually.
type Store interface {
	SavePage(*model.Page) (*model.Page, error)
	RemoveAll() error
	GetPage(slug string) (*model.Page, error)
	GetMenu() ([]*model.PageFragment, error)
	Search(q string) ([]*model.PageFragment, error)
}
