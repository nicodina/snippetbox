package models

import (
	"errors"
	"time"
)

// ErrNoRecord is thrown when no record is found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet represents the data model of table snippets
type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}