package main

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lordmoma/blog-aggregator/internal/database"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	// Description is optional
	Description *string `json:"description"`
	Url         string  `json:"url"`
	// PublishedAt is optional
	PublishedAt *time.Time `json:"published_at"`
}

func (apiCfg *apiConfig) getPostsHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	params := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  10,
	}
	posts, err := apiCfg.DB.GetPostsByUser(context.Background(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get posts")
		return
	}
	result := make([]Post, len(posts))
	for i, post := range posts {
		params := Post{
			ID:        uuid.New(),
			FeedID:    post.FeedID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     post.Title,
			Description: func() *string {
				if post.Description.Valid {
					return &post.Description.String
				}
				return nil
			}(),
			Url: post.Url,
			PublishedAt: func() *time.Time {
				if post.PublishedAt.Valid {
					return &post.PublishedAt.Time
				}
				return nil
			}(),
		}
		result[i] = params
	}
	respondWithJSON(w, http.StatusOK, result)
}
