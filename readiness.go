package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
