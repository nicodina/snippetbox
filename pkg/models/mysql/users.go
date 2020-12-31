package mysql

import (
	"database/sql"

	"github.com/nicodina/snippetbox/pkg/models"
)

// UsersService allows to perform actions on the database
type UsersService struct {
	DB *sql.DB
}

// Insert writes a new user in the database
func (s *UsersService) Insert(name, email, password string) error {
	return nil
}

// Authenticate checks if the user exists in the database
func (s *UsersService) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get retrieves a specific user from the database
func (s *UsersService) Get(id int) (*models.User, error) {
	return nil, nil
}
