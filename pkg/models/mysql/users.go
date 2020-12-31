package mysql

import (
	"database/sql"

	"github.com/nicodina/snippetbox/pkg/models"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// UsersService allows to perform actions on the database
type UsersService struct {
	DB *sql.DB
}

// Insert writes a new user in the database
func (s *UsersService) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
			 VALUES(?, ?, ?, UTC_TIMESTAMP())`
	
	_, err = s.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return models.ErrDuplicateEmail
			}
		}
	}
	return err
}

// Authenticate checks if the user exists in the database
func (s *UsersService) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	query := "SELECT id, hashed_password FROM users WHERE email = ?"
	row := s.DB.QueryRow(query, email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

// Get retrieves a specific user from the database
func (s *UsersService) Get(id int) (*models.User, error) {
	return nil, nil
}
