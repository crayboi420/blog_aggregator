package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

func fetchData(Url string) (*RSSFeed, error) {
	net := &http.Client{Timeout: time.Second * 10}
	r, err := net.Get(Url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}
	newFeed := RSSFeed{}
	err = xml.Unmarshal(body, &newFeed)
	if err != nil {
		return nil, err
	}
	return &newFeed, nil
}
