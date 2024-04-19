package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ivanpatera/twclone/internal/domain/usecases"
	"github.com/ivanpatera/twclone/pkg/auth"
)

type TweetHandler struct {
	TweetUsecase usecases.TweetUseCase
}

type NewTweetRequest struct {
	Text string `json:"text"`
}

type NewTweetResponse struct {
	Message string `json:"message"`
	Tweet  TweetDetailsResponse `json:"tweet"`
}

type TweetDetailsResponse struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	Text string `json:"text"`
	Ts int64 `json:"ts"`
}

func (handler TweetHandler) NewTweetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "text" param from the request body
	body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

	// Unmarshal the JSON data into the struct
    var req NewTweetRequest
    if err := json.Unmarshal(body, &req); err != nil {
        http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
        return
    }

	// Get the user from the request headers
	// user := r.Header.Get("x-user-id")
	user := r.Context().Value(auth.UserIDKey)
	// Perform type assertion to convert to string
	uniqueIdentifier, err := uuid.Parse(user.(string))
	if err != nil {
		http.Error(w, "INVALID_USER", http.StatusBadRequest)
		return
	}
	tw, err := handler.TweetUsecase.NewTweetUseCase(req.Text, uniqueIdentifier)
	if err != nil {
		if err.Error() == "TWEET_TOO_LONG" {
			http.Error(w, "TWEET_TOO_LONG", http.StatusBadRequest)
		} else {
			http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		}
		return
	}
	json, err := json.Marshal(
		NewTweetResponse{
			Message: "Tweet post was successful", 
			Tweet: TweetDetailsResponse{
				ID: tw.ID.String(), 
				UserId: tw.UserId.String(), 
				Text: tw.Text, 
				Ts: tw.Ts,
			},
		})
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return

	}
	w.Write(json)
}

type GetTweetResponse struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	Text string `json:"text"`
	Ts int64 `json:"ts"`
}

func (handler TweetHandler) GetTweetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "id" param from the request URL
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Missing tweet ID", http.StatusBadRequest)
		return
	}

	// Get the tweet by its ID
	uniqueIdentifier, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}
	tw, err := handler.TweetUsecase.GetTweetById(uniqueIdentifier)
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(GetTweetResponse{ID: tw.ID.String(), UserId: tw.UserId.String(), Text: tw.Text, Ts: tw.Ts})
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return

	}
	w.Write(json)
}
