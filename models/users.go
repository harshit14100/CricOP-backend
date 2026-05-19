package models

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	PhoneNo  string    `json:"phone_no"`
	Password string    `json:"-"`
}
