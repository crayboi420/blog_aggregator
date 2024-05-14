package main

import (
	"encoding/json"
	"github.com/crayboi420/blog_aggregator/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerFeedsPost(w http.ResponseWriter, r *http.Request, user database.User) {
	type inpt struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	inp := inpt{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inp)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
	}

	new_feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      inp.Name,
		Url:       inp.Url,
		UserID:    user.ID,
	}
	resp1, err := cfg.DB.CreateFeed(
		r.Context(),
		new_feed,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed: "+err.Error())
		return
	}

	type ret struct {
		Feed       Feed                `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}

	new_follow := database.CreateFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    resp1.ID,
		UserID:    user.ID,
	}

	resp2, err := cfg.DB.CreateFollow(r.Context(), new_follow)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow: "+err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ret{Feed: databaseFeedtoFeed(resp1), FeedFollow: resp2})
}

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	dbfeeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't find feeds: "+err.Error())
		return
	}
	feeds := make([]Feed, 0)

	for _, dbfeed := range dbfeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbfeed))
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
