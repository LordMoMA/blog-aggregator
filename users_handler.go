package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lordmoma/blog-aggregator/internal/database"
)

type userRequest struct {
	Name string `json:"name"`
}

func (apiCfg *apiConfig) userHandler(w http.ResponseWriter, r *http.Request) {
	var req userRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      req.Name,
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondwithJSON(w, http.StatusOK, user)
}
