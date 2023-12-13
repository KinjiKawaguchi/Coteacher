package models

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID
	Name      string
	TeacherID uuid.UUID
}

type ClassInvitationCode struct {
	ID             uuid.UUID
	ClassID        uuid.UUID
	InvitationCode string
	ExpirationDate time.Time
	IsActive       bool
}
