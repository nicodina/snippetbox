package mysql

import (
	"database/sql"

	"github.com/nicodina/snippetbox/pkg/models"
)

// SnippetService encapsulate a connection pool
type SnippetService struct {
	DB *sql.DB
}

// Insert will insert a new snippet into the database.
func (s *SnippetService) Insert(title, content, expires string) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires)
			 VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := s.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get will return a specific snippet based on its id.
func (s *SnippetService) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest will return the 10 most recently created snippets.
func (s *SnippetService) Latest() ([]*models.Snippet, error) {
	return nil, nil
}