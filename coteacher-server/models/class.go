package models

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
)


type Class struct {
	ID uuid.UUID
	Name string
	TeacherID uuid.UUID
}

type ClassInvitationCode struct {
	ID             uuid.UUID
	ClassID        uuid.UUID
	InvitationCode string
	ExpirationDate date.Date
	IsActive bool
}