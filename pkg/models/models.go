package models

import (
	"errors"
	"time"
)

// ErrNoRecord is thrown when no record is found
var (
	ErrNoRecord = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail = errors.New("models: duplicated email")
)

// Snippet represents the data model of table snippets
type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}
//User represents the data model of table users
type User struct {
	ID int
	Name string
	Email string
	HashedPassword []byte
	Created time.Time
}