package entities

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID
	Username string
	FollowerCount uint
}