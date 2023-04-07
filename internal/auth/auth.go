package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func GetApiKey(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	apiKey := strings.TrimPrefix(authHeader, "ApiKey ")
	if apiKey == "" {
		return "", errors.New("no api key included")
	}
	return apiKey, nil
}
