package email

import (
	"database/sql"
	"log"
)

type UserFetcher interface {
	FetchUserById(userId int) (string, error)
}

type UserEmailRepository struct {
	db *sql.DB
}

// NewUserEmailRepository creates a new instance of UserEmailRepository.
func NewUserEmailRepository(db *sql.DB) *UserEmailRepository {
	return &UserEmailRepository{db: db}
}

// Implements the UserFetcher interface
func (u *UserEmailRepository) FetchUserById(userId int) (string, error) {
    var userEmail string
    query := "SELECT email FROM users WHERE user_id = ?"

    err := u.db.QueryRow(query, userId).Scan(&userEmail)
	if err != nil {
		log.Println(err)
		return userEmail, err
	}
	return userEmail, nil
}
