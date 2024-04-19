package sql

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ivanpatera/twclone/pkg/sql"
)

type FollowSqlAdapter struct{}

func (f *FollowSqlAdapter) Follow(followerId, followeeId uuid.UUID) error {
	qry := `INSERT INTO followings (follower_id, followee_id) VALUES ($1, $2)`
	res, err := sql.InsertRow(qry, followerId.String(), followeeId.String())
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("ALREADY_FOLLOWING_USER")
	}
	return nil
}

func (f *FollowSqlAdapter) Unfollow(followerId, followeeId uuid.UUID) error {
	qry := `DELETE FROM followings WHERE follower_id = $1 AND followee_id = $2`
	res, err := sql.DeleteRow(qry, followerId.String(), followeeId.String())
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("ROWS_AFFECTED_NOT_SUPPORTED_BY_DRIVER")
	}
	if rowsAffected == 0 {
		return errors.New("NOT_FOLLOWING_USER")
	}
	return nil
}
