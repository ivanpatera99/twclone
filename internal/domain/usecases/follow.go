package usecases

import (
	"github.com/google/uuid"
	"github.com/ivanpatera/twclone/internal/domain/ports"
)

type FollowUseCase struct {
	Repo ports.FollowRepository
}

func (usecase FollowUseCase) Follow(followerId, followedId string) error {
	follower, err := uuid.Parse(followerId)
	if err != nil {
		return err
	}
	followee, err := uuid.Parse(followedId)
	if err != nil {
		return err
	}
	return usecase.Repo.Follow(follower, followee)
}

func (usecase FollowUseCase) Unfollow(followerId, followedId string) error {
	follower, err := uuid.Parse(followerId)
	if err != nil {
		return err
	}
	followee, err := uuid.Parse(followedId)
	if err != nil {
		return err
	}
	err = usecase.Repo.Unfollow(follower, followee)
	if err != nil {
		return err
	}
	return nil
}