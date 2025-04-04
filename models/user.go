package models

import (
	"errors"
	"golang-web-app/config"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func AuthenticateUser(email, password string) (*User, error) {
	var user User
	query := "SELECT id, email FROM users WHERE email = ? AND password = ?"
	err := config.DB.QueryRow(query, email, password).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &user, nil
}
