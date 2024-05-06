package main

import (
	"database/sql"
	"fmt"
	"github.com/crayboi420/blog_aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")

	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	cfg := apiConfig{DB: dbQueries}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/readiness", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)
	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)

	cors := middlewareCORS(mux)
	serv := &http.Server{
		Addr:    ":" + port,
		Handler: cors,
	}
	fmt.Printf("Listening and serving on Port %v\n", port)
	log.Fatal(serv.ListenAndServe())
}

func middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
