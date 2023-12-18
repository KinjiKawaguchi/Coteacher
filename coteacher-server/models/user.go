package models

import "github.com/google/uuid"

type User struct {
	ID 	 uuid.UUID
	Email    string
	Name     string
	UserType string
}