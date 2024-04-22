package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ivanpatera/twclone/internal/domain/usecases"
	"github.com/ivanpatera/twclone/pkg/auth"
)

type TimelineHandler struct {
	TimelineUsecase usecases.TimelineUseCase
}

type GetTimelineFollowingResponse struct {
	Timeline []TimelineTweetResponse `json:"timeline"`
	Pagination PaginationResponse     `json:"pagination"`
}

type TimelineTweetResponse struct {
	ID       string `json:"id"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Ts       int64  `json:"ts"`
}

type PaginationResponse struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

func (handler TimelineHandler) GetTimelineFollowingHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth.UserIDKey).(string)
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "INVALID_PARAMETER", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		http.Error(w, "INVALID_PARAMETER", http.StatusBadRequest)
		return
	}
	tl, err := handler.TimelineUsecase.GetTimelineFollowing(userId, limit, offset)
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
	response := GetTimelineFollowingResponse{
		Timeline: make([]TimelineTweetResponse, len(tl.Tweets)),
		Pagination: PaginationResponse{
			Limit: limit,
			Offset: offset,
		},
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
