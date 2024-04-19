package usecases_test

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/ivanpatera/twclone/internal/domain/entities"
	"github.com/ivanpatera/twclone/internal/domain/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockTweetRepository struct {
    mock.Mock
}

func (m *MockTweetRepository) PostTweet(userId uuid.UUID, text string) (*entities.Tweet, error) {
    args := m.Called(userId, text)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*entities.Tweet), args.Error(1)
}

func (m *MockTweetRepository) GetTweetById(id uuid.UUID) (*entities.Tweet, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*entities.Tweet), args.Error(1)
}

func TestNewTweetUseCase(t *testing.T) {
    mockRepo := new(MockTweetRepository)
    usecase := usecases.TweetUseCase{Repo: mockRepo}
    userId := uuid.New()
    text := "Hello, world!"

    mockRepo.On("PostTweet", userId, text).Return(&entities.Tweet{ID: uuid.New(), UserId: userId, Text: text}, nil)

    tweet, err := usecase.NewTweetUseCase(text, userId)
    require.NoError(t, err)
    require.Equal(t, text, tweet.Text)
    require.Equal(t, userId, tweet.UserId)

    mockRepo.AssertExpectations(t)
}

func TestNewTweetUseCase_WithMoreThan280Chars(t *testing.T) {
    usecase := usecases.TweetUseCase{Repo: new(MockTweetRepository)}
    userId := uuid.New()
    text := strings.Repeat("a", 281) // This creates a string with 281 characters

    tweet, err := usecase.NewTweetUseCase(text, userId)
    require.Error(t, err)
    require.Equal(t, entities.Tweet{}, tweet)
    require.Equal(t, "TWEET_TOO_LONG", err.Error())
}

func TestGetTweetById(t *testing.T) {
    mockRepo := new(MockTweetRepository)
    usecase := usecases.TweetUseCase{Repo: mockRepo}
    tweetId := uuid.New()
    userId := uuid.New()
    text := "Hello, world!"

    mockRepo.On("GetTweetById", tweetId).Return(&entities.Tweet{ID: tweetId, UserId: userId, Text: text}, nil)

    tweet, err := usecase.GetTweetById(tweetId)
    require.NoError(t, err)
    require.Equal(t, tweetId, tweet.ID)
    require.Equal(t, userId, tweet.UserId)
    require.Equal(t, text, tweet.Text)

    mockRepo.AssertExpectations(t)
}