package usecases

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ivanpatera/twclone/internal/domain/entities"
	"github.com/ivanpatera/twclone/internal/domain/ports"
)

type TweetUseCase struct {
	Repo ports.TweetRepository
}

func (usecase TweetUseCase) NewTweetUseCase(text string, userId uuid.UUID) (entities.Tweet, error) {
	if len(text) > 280 {
		return entities.Tweet{}, errors.New("TWEET_TOO_LONG")
	}
	tw, err := usecase.Repo.PostTweet(userId, text)
	if err != nil {
		return entities.Tweet{}, err
	}
	return *tw, nil
}

func (usecase TweetUseCase) GetTweetById(id uuid.UUID) (entities.Tweet, error) {
	tw, err := usecase.Repo.GetTweetById(id)
	if err != nil {
		return entities.Tweet{}, err
	}
	return *tw, nil
}
