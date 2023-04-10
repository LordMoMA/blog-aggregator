package main

import (
	"context"
	"net/http"

	"github.com/lordmoma/blog-aggregator/internal/database"
)

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
	respondWithJSON(w, http.StatusOK, posts)
}
