package models

import "github.com/google/uuid"

type User struct {
	ID 	 uuid.UUID
	Name    string
	Email     string
	UserType string
}