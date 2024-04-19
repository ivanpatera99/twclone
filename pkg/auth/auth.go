package auth

import (
	"errors"

	"github.com/ivanpatera/twclone/pkg/sql"

	"github.com/google/uuid"
)

type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	UsernameKey contextKey = "username"
)

type User struct {
	ID       string
	Username string
}

func GetUser(userId string, username string) (User, error) {
	var user User
	row, err := sql.QueryRow("SELECT id AS 'ID', username AS 'Username' FROM users WHERE id = ?", userId)
	if err != nil {
		return User{}, err
	}
	err = row.Scan(&user.ID, &user.Username)
	// if the user does not exist, create a new one with the provided username
	// additionaly, if username is empty, request fails
	if err != nil && user.Username != "" {
		id := uuid.New().String()
		_, err := sql.InsertRow("INSERT INTO users (id, username, follower_count) VALUES (?, ?, 0)", id, username)
		if err != nil {
			return User{}, err
		}
		return User{ID: id, Username: username}, nil
	} 
	return User{}, errors.New("INVALID_USER")
}
