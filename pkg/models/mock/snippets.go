package mock

import (
	"time"

	"github.com/nicodina/snippetbox/pkg/models"
)

var mockedSnippet = &models.Snippet{
	ID: 1,
	Title: "A mocked snippet",
	Content: "A mocked snippet content",
	Created: time.Now(),
	Expires: time.Now(),
}

// SnippetService is a mock of a snippet service
type SnippetService struct {}

// Insert fakes a write in a database
func (s *SnippetService) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

// Get fakes a read from a database based on the id
func (s *SnippetService) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockedSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest fakes a read of last inserted snippets
func (s *SnippetService) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockedSnippet}, nil
}