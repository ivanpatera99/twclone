package sql

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivanpatera/twclone/internal/domain/entities"
	"github.com/ivanpatera/twclone/pkg/sql"
)

type TweetSqlAdapter struct {
}

func (a *TweetSqlAdapter) GetTweetById(id uuid.UUID) (*entities.Tweet, error) {
	// Implementation for getting a tweet by its ID from the SQL database
	query := "SELECT * FROM tweets WHERE id = ?"
	row, err := sql.QueryRow(query, id.String())
	if err != nil {
		return nil, err
	}
	tw := entities.Tweet{}
	row.Scan(&tw.ID, &tw.UserId, &tw.Text, &tw.Ts)
	return &tw, nil
}

func (a *TweetSqlAdapter) PostTweet(userId uuid.UUID, text string) (*entities.Tweet, error) {
	// Implementation for posting a tweet to the SQL database
	id := uuid.New()
	ts := time.Now().Unix()
	query := "INSERT INTO tweets (id, user_id, text, ts) VALUES (?, ?, ?, ?)"
	_, err := sql.InsertRow(query, id.String(), userId.String(), text, ts)
	if err != nil {
		return nil, err
	}
	return &entities.Tweet{ID: id, UserId: userId, Text: text, Ts: ts}, nil
}
