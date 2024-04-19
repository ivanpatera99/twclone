package ports

import (
	"github.com/ivanpatera/twclone/internal/domain/entities"
	"github.com/google/uuid"
)

type TweetRepository interface {
	GetTweetById(id uuid.UUID) (*entities.Tweet, error);
	PostTweet(userId uuid.UUID, text string) (*entities.Tweet, error);
}