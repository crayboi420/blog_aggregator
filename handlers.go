package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/crayboi420/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{"ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type userStruct struct {
		Name string `json:"name"`
	}
	usr := userStruct{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Coulnd't decode parameters")
	}
	name := usr.Name
	uuid := uuid.New()
	user := database.CreateUserParams{
		Name:      name,
		ID:        uuid,
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
