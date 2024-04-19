package sql

import (
	"errors"
	"github.com/ivanpatera/twclone/internal/domain/entities"
	"github.com/ivanpatera/twclone/pkg/sql"
)

type TimelineSqlAdapter struct{}

func (a TimelineSqlAdapter) GetTimelineFollowing(userId string) (entities.Timeline, error) {
	qry := `
		SELECT t.id, t.text, t.ts, u.id, u.username
		FROM tweets t
		JOIN followings f ON f.follower_id = ? AND f.followee_id = t.user_id
		JOIN users u ON t.user_id = u.id
	`
	rows, err := sql.QueryRows(qry, userId)
	if err != nil {
		return entities.Timeline{}, err
	}
	defer rows.Close()
	
	var timeline = entities.Timeline{Tweets: []entities.TweetInTimeline{}}
	for rows.Next() {
		var tw entities.TweetInTimeline
		err = rows.Scan(&tw.ID, &tw.Text, &tw.Ts, &tw.UserId, &tw.Username)
		if err != nil {
			return entities.Timeline{}, err
		}
		timeline.Tweets = append(timeline.Tweets, tw)
	}
	if len(timeline.Tweets) == 0{
		return entities.Timeline{}, errors.New("NO_TIMELINE_FOR_USER")
	}
	return timeline, nil
}
