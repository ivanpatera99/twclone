package ports

import (
	"github.com/google/uuid"
)

type FollowRepository interface {
	Follow(followerId, followedId uuid.UUID) error
	Unfollow(followerId, followedId uuid.UUID) error
}