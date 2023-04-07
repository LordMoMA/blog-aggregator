package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lordmoma/blog-aggregator/internal/auth"
	"github.com/lordmoma/blog-aggregator/internal/database"
)

type feedRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type feedResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func (apiCfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request) {
	var req feedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	apiKey, err := auth.GetApiKey(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get api key")
		return
	}
	user, err := apiCfg.DB.GetUserbyApiKey(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
		Url:       req.URL,
		UserID:    user.ID,
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), params)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusOK, feedResponse{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		URL:       feed.Url,
		UserID:    feed.UserID,
	})
}
