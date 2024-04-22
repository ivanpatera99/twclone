package usecases

import (
	"github.com/ivanpatera/twclone/internal/domain/entities"
	"github.com/ivanpatera/twclone/internal/domain/ports"
)

type TimelineUseCase struct {
	Repo ports.TimelineRepository
	// Documents ports.DocumentRepository
	// Graphs ports.GraphRepository
}

func (usecase TimelineUseCase) GetTimelineFollowing(userId string, limit, offset int) (entities.Timeline, error) {
	// query the graph database for the graph of users that compose our timeline
	// usecase.Graphs.GetTimelineGraph(userId)

	// query the document database for the tweets of the users in the graph
	// usecase.Documents.GetTweetsFromGraph(usersInGraph)

	// in an internal function of this package, mix the tweets to give them an organic order
	// this mixing function can be a usefull way of alterate tweets from different origins
	// once the app complexity grows
	return usecase.Repo.GetTimelineFollowing(userId, limit, offset)
}
