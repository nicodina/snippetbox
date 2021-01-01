package mock

import (
	"time"

	"github.com/nicodina/snippetbox/pkg/models"
)

var mockedUser = &models.User{
	ID: 1,
	Name: "Zlatan Ibrahimovic",
	Email: "zlatan@ibra.com",
	HashedPassword: []byte("password"),
	Created: time.Now(),
}

// UsersService is a fake service on users table
type UsersService struct{}

// Insert fakes a write in a database
func (s *UsersService) Insert(name, email, password string) error {
	switch email {
	case "dup@email.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// Authenticate checks if the user exists in the database
func (s *UsersService) Authenticate(email, password string) (int, error) {
	switch email {
	case "zlatan@ibra.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

// Get retrieves a specific user from the database
func (s *UsersService) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockedUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}