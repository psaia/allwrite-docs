package postgres

import (
	"database/sql"
	"fmt"

	"github.com/LevInteractive/allwrite-docs/model"

	// Because postgres
	_ "github.com/lib/pq"
)

// Store implements the Store interface.
type Store struct {
	driver *sql.DB
}

// Init connection with postgres.
func Init(user string, pass string, host string, db string) (*Store, error) {
	driver, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, db),
	)
	if err != nil {
		return &Store{}, err
	}
	if err = driver.Ping(); err != nil {
		return &Store{}, err
	}

	return &Store{driver: driver}, nil
}

// SavePage saves a page.
func (p *Store) SavePage(*model.Page) (*model.Page, error) {
	return nil, nil
}

// RemoveAll removes all pages from postgres.
func (p *Store) RemoveAll() error {
	return nil
}

// GetPage saves a page.
func (p *Store) GetPage(slug string) (*model.Page, error) {
	return nil, nil
}

// GetMenu retrieves the full menu tree.
func (p *Store) GetMenu() ([]*model.PageFragment, error) {
	return nil, nil
}

// Search searches a page.
func (p *Store) Search(q string) ([]*model.PageFragment, error) {
	return nil, nil
}
