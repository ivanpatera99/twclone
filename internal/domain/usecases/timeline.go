package usecases

import (
	"github.com/ivanpatera/twclone/internal/domain/entities"
	"github.com/ivanpatera/twclone/internal/domain/ports"
)

type TimelineUseCase struct {
	Repo ports.TimelineRepository
}

func (usecase TimelineUseCase) GetTimelineFollowing(userId string, limit, offset int) (entities.Timeline, error) {
	return usecase.Repo.GetTimelineFollowing(userId, limit, offset)
}
