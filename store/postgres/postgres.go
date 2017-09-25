package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/LevInteractive/allwrite-docs/model"

	// Because postgres
	_ "github.com/lib/pq"
)

// Store implements the Store interface.
type Store struct {
	driver *sql.DB
}

// SavePages saves a page. This will preform an upsert on the page record.
// It's assumed that the slug is already unique.
//
// The upsert here is also completely unncessary because we are removing all
// rows anytime it's updated. Just keeping it because it doesn't hurt.
func (p *Store) SavePages(pages model.Pages) (model.Pages, error) {
	tx, err := p.driver.Begin()
	if err != nil {
		return pages, err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO pages (doc_id, type, title, slug, md, placement)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (doc_id)
		DO UPDATE set (type, title, slug, md, placement) = ($2, $3, $4, $5, $6)
		WHERE pages.doc_id = $1
	`)

	if err != nil {
		return pages, err
	}

	defer stmt.Close()

	var statementError error

	for _, page := range pages {
		if _, err := stmt.Exec(
			page.DocID,
			page.Type,
			page.Name,
			page.Slug,
			page.Md,
			page.Order,
		); err != nil {
			tx.Rollback()
			statementError = err
			break
		}
	}

	if statementError != nil {
		return pages, statementError
	}

	tx.Commit()
	return pages, nil
}

// RemoveAll removes all pages from postgres.
func (p *Store) RemoveAll() error {
	if _, err := p.driver.Exec(`DELETE FROM pages`); err != nil {
		return err
	}
	return nil
}

// GetPage saves a page.
func (p *Store) GetPage(slug string) (*model.Page, error) {
	stmt := `
	SELECT doc_id, type, title, slug, md, placement
	FROM pages
	WHERE slug = $1 AND type = "file"`
	var page model.Page

	err := p.driver.QueryRow(stmt, slug).Scan(
		&page.DocID,
		&page.Type,
		&page.Name,
		&page.Slug,
		&page.Md,
		&page.Order,
	)

	switch err {
	case sql.ErrNoRows:
		return &page, errors.New("not found")
	case nil:
		return &page, nil
	default:
		return &page, err
	}
}

// GetMenu retrieves the full menu tree.
func (p *Store) GetMenu() ([]*model.PageFragment, error) {
	return nil, nil
}

// Search searches a page.
func (p *Store) Search(q string) ([]*model.PageFragment, error) {
	return nil, nil
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
