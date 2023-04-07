package main

import (
	"net/http"
	"strings"

	"github.com/lordmoma/blog-aggregator/internal/database"
)

// authedHandler is a custom type for handlers that require authentication.
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "No auth header")
			return
		}
		apiKey := strings.TrimPrefix(authHeader, "ApiKey ")
		if apiKey == "" {
			respondWithError(w, http.StatusUnauthorized, "No api key")
			return
		}

		user, err := apiCfg.DB.GetUserbyApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get users")
			return
		}
		handler(w, r, user)
	}
}
