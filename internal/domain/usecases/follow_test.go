package usecases_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ivanpatera/twclone/internal/domain/usecases"
	"github.com/stretchr/testify/mock"
)

type MockFollowRepository struct {
    mock.Mock
}

func (m *MockFollowRepository) Follow(follower, followee uuid.UUID) error {
    args := m.Called(follower, followee)
    return args.Error(0)
}

func (m *MockFollowRepository) Unfollow(follower, followee uuid.UUID) error {
	args := m.Called(follower, followee)
	return args.Error(0)
}

func TestFollowUseCase_Follow(t *testing.T) {
    mockRepo := new(MockFollowRepository)
    u := usecases.FollowUseCase{
        Repo: mockRepo,
    }

    followerId := uuid.New().String()
    followedId := uuid.New().String()

    mockRepo.On("Follow", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("uuid.UUID")).Return(nil)

    err := u.Follow(followerId, followedId)

    mockRepo.AssertExpectations(t)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}

func TestFollowUseCase_Unfollow(t *testing.T) {
    mockRepo := new(MockFollowRepository)
    u := usecases.FollowUseCase{
        Repo: mockRepo,
    }

    followerId := uuid.New().String()
    followedId := uuid.New().String()

    mockRepo.On("Unfollow", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("uuid.UUID")).Return(nil)

    err := u.Unfollow(followerId, followedId)

    mockRepo.AssertExpectations(t)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}