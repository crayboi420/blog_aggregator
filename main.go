package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"os"

	"github.com/kanavj/blog_aggregator/internal/database"
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"net/http"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")

	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	cfg := apiConfig{DB: dbQueries}

	go cfg.continousFetching(3, 10*time.Second)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/readiness", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(cfg.handlerUsersGet))

	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.handlerFeedsPost))
	mux.HandleFunc("GET /v1/feeds", cfg.handlerFeedsGet)

	mux.HandleFunc("POST /v1/feed_follows", cfg.middlewareAuth(cfg.handlerFeedFollowsPost))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", cfg.handlerFeedFollowsDelete)
	mux.HandleFunc("GET /v1/feed_follows", cfg.middlewareAuth(cfg.handlerFeedFollowsGet))

	mux.HandleFunc("GET /v1/posts", cfg.middlewareAuth(cfg.handlerPostsGet))

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
