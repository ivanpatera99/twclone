package ports

import "github.com/ivanpatera/twclone/internal/domain/entities"

type TimelineRepository interface {
	GetTimelineFollowing(userId string, limit, offset int) (entities.Timeline, error)
	// GetTimelineDiscover(userId string) (entities.Timeline, error)
}