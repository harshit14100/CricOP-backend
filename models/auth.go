package models

import "github.com/google/uuid"

type SignupRequest struct {
	Name     string `json:"name"`
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	PhoneNo  string    `json:"phone_no"`
	Password string    `json:"-"`
}

type PasswordRequest struct {
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}
