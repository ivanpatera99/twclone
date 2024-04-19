package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanpatera/twclone/internal/domain/usecases"
	"github.com/ivanpatera/twclone/pkg/auth"
)

type FollowHandler struct {
	FollowUsecase usecases.FollowUseCase
}

type HandlerResponse struct {
	Message string `json:"message"`
}

func (handler FollowHandler) FollowHandler(w http.ResponseWriter, r *http.Request) {
	follower := r.Context().Value(auth.UserIDKey).(string)
	userId := mux.Vars(r)["user_id"]
	if userId == "" {
		http.Error(w, "FOLLOWEE_EMPTY", http.StatusBadRequest)
		return
	}
	err := handler.FollowUsecase.Follow(follower, userId)
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(HandlerResponse{Message: "Follow success"})
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return

	}
	w.Write(json)
}

func (handler FollowHandler) UnfollowHandler(w http.ResponseWriter, r *http.Request) {
	follower := r.Context().Value(auth.UserIDKey).(string)
	userId := mux.Vars(r)["user_id"]
	if userId == "" {
		http.Error(w, "FOLLOWEE_EMPTY", http.StatusBadRequest)
		return
	}
	err := handler.FollowUsecase.Unfollow(follower, userId)
	if err != nil {
		if err.Error() == "FOLLOW_NOT_FOUND" {
			http.Error(w, "FOLLOW_NOT_FOUND", http.StatusBadRequest)
			return
		}
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(HandlerResponse{Message: "Follow success"})
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return

	}

	w.Write(json)
}
