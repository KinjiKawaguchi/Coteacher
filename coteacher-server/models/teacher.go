package models

import "github.com/google/uuid"

type Teacher struct {
	ID uuid.UUID
	Email     string
	Name      string
}