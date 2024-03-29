package models

import (
	"database/sql"
	"fmt"
	"strings"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type NewUser struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GeneratedFromPassword([]byte(password), bcrypt, DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Create User Failed: %w", err)
	}
	passwordHash := string(hashedBytes)
	return nil, nil
}
