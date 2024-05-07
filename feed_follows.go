package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/crayboi420/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowsPost(w http.ResponseWriter, r *http.Request, user database.User) {
	type inpt struct{
		FeedId uuid.UUID `json:"feed_id"`
	}
	inp := inpt{}
	decoder := json.NewDecoder(r.Body)
	err:= decoder.Decode(&inp)
	if err!=nil{
		respondWithError(w,http.StatusBadRequest,"Couldn't decode parameters")
		return
	}

	new_follow:= database.CreateFollowParams{
		ID: uuid.New(),
		FeedID: inp.FeedId,
		UserID: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	resp,err := cfg.DB.CreateFollow(r.Context(),new_follow)
	if err!=nil{
		respondWithError(w,http.StatusInternalServerError,"Couldn't create feed follow: "+err.Error())
		return
	}
	respondWithJSON(w,http.StatusOK,resp)
}

func (cfg *apiConfig) handlerFeedFollowsDelete(w http.ResponseWriter,r *http.Request){
	followID,err := uuid.Parse(r.PathValue("feedFollowID"))
	if err!=nil{
		respondWithError(w, http.StatusBadRequest, "Not a valid feed ID: "+err.Error())
		return 
	}

	err = cfg.DB.DeleteFollow(r.Context(),followID)
	if err!= nil{
		respondWithError(w,http.StatusInternalServerError,"Couldn't delete follow: "+err.Error())
		return 
	}
	respondWithJSON(w,http.StatusOK,"{}")
}

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User){
	resp,err := cfg.DB.GetFollows(r.Context(),user.ID)
	if err!=nil{
		respondWithError(w, http.StatusInternalServerError,"Couldn't retrieve follows :"+err.Error())
		return
	}
	respondWithJSON(w,http.StatusOK,resp)
}