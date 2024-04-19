package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanpatera/twclone/internal/adapters/sql"
	"github.com/ivanpatera/twclone/internal/app/http/handlers"
	"github.com/ivanpatera/twclone/internal/domain/usecases"
	"github.com/ivanpatera/twclone/pkg/middleware"
)

func App() {
	// Tweet deps
	tweetRepo := sql.TweetSqlAdapter{}
	tweetUsecase := usecases.TweetUseCase{Repo: &tweetRepo}
	twHandler := handlers.TweetHandler{TweetUsecase: tweetUsecase}

	// Follow deps
	followRepo := sql.FollowSqlAdapter{}
	followUsecase := usecases.FollowUseCase{Repo: &followRepo}
	fHandler := handlers.FollowHandler{FollowUsecase: followUsecase}

	// Timeline deps
	timelineRepo := sql.TimelineSqlAdapter{}
	timelineUsecase := usecases.TimelineUseCase{Repo: &timelineRepo}
	tlHandler := handlers.TimelineHandler{TimelineUsecase: timelineUsecase} 


	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	// Tweet routes
	r.HandleFunc("/tweet", twHandler.NewTweetHandler).Methods("POST")
	r.HandleFunc("/tweet/{id}", twHandler.GetTweetHandler).Methods("GET")

	// Follow routes
	r.HandleFunc("/follow/{user_id}", fHandler.FollowHandler).Methods("POST")
	r.HandleFunc("/unfollow/{user_id}", fHandler.UnfollowHandler).Methods("POST")

	// Timeline routes
	r.HandleFunc("/timeline", tlHandler.GetTimelineFollowingHandler).Methods("GET")

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
