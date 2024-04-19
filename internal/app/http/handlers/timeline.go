package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ivanpatera/twclone/internal/domain/usecases"
	"github.com/ivanpatera/twclone/pkg/auth"
)

type TimelineHandler struct {
	TimelineUsecase usecases.TimelineUseCase
}

type GetTimelineFollowingResponse struct {
	Timeline []TimelineTweetResponse `json:"timeline"`
}

type TimelineTweetResponse struct {
	ID       string `json:"id"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Ts       int64  `json:"ts"`
}

func (handler TimelineHandler) GetTimelineFollowingHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth.UserIDKey).(string)
	tl, err := handler.TimelineUsecase.GetTimelineFollowing(userId)
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
	response := GetTimelineFollowingResponse{
		Timeline: make([]TimelineTweetResponse, len(tl.Tweets)),
	}
	for i, tweet := range tl.Tweets {
		response.Timeline[i] = TimelineTweetResponse{
			ID:       tweet.ID,
			UserId:   tweet.UserId,
			Username: tweet.Username,
			Text:     tweet.Text,
			Ts:       tweet.Ts,
		}
	}
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
