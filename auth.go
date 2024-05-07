package main

import (
	"net/http"
	"strings"
	"github.com/crayboi420/blog_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter,*http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		apiKey := strings.TrimPrefix(r.Header.Get("Authorization"), "ApiKey ")
		
		retr, err := cfg.DB.GetUserApi(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find ApiKey")
			return
		}
		handler(w,r,retr)
	}
}