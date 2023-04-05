package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/lordmoma/blog-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	port := os.Getenv("PORT")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r2 := chi.NewRouter()
	r2.Get("/readiness", readinessHandler)
	r2.Get("/err", errHandler)

	r2.Post("/users", userHandler)

	r.Mount("/v1", r2)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Serving files on port: %s\n", port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"error": msg})
}
