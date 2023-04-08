package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
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

type feedFollowRequest struct {
	FeedID uuid.UUID `json:"feed_id"`
}

type feedFollowResponse struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (apiCfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	var req feedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
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

func (apiCfg *apiConfig) getFeedHandler(w http.ResponseWriter, r *http.Request) {
	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, feed)
}

func (apiCfg *apiConfig) createFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	var req feedFollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    req.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollowResponse{
		ID:        feedFollow.ID,
		FeedID:    feedFollow.FeedID,
		UserID:    feedFollow.UserID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	})
}

func (apiCfg *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	idString := chi.URLParam(r, "feedFollowID")
	fmt.Println(idString)
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse feed_follow id")
		return
	}

	if err := apiCfg.DB.DeleteFeedFollow(r.Context(), id); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

// getFeedFollowHandler
func (apiCfg *apiConfig) getFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollow)
}
