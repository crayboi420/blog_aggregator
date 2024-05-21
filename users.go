package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kanavj/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type userStruct struct {
		Name string `json:"name"`
	}
	usr := userStruct{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
	}

	user := database.CreateUserParams{
		Name:      usr.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	returned, err := cfg.DB.CreateUser(r.Context(), user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't add to database : "+err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, returned)
}

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request,user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}
