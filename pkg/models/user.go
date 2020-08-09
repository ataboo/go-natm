package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID
	Name         string
	Email        string
	JWTToken     string
	RefreshToken string
}
