package models

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type UserDTO struct {
	Email    string
	Password string
}

func NewUser(email, passwordHash string) User {
	return User{
		Email:        email,
		PasswordHash: passwordHash,
	}
}
