package entities

import "github.com/google/uuid"

type Tweet struct {
	ID uuid.UUID
	UserId uuid.UUID
	Text string
	Ts int64
}

