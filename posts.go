package main

import (
	"github.com/kanavj/blog_aggregator/internal/database"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitstr := r.URL.Query().Get("limit")
	var limit int32
	if limitstr == "" {
		limit = 5
	} else {
		readlimit, err := strconv.Atoi(limitstr)
		limit = int32(readlimit)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "limit not integer")
			return
		}
	}
	DBposts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't read posts")
		return
	}
	posts := make([]Post, 0)
	for _, dbPost := range DBposts {
		posts = append(posts, databasePosttoPost(dbPost))
	}
	respondWithJSON(w, http.StatusOK, posts)
}
