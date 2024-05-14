package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/crayboi420/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func fetchData(Url string) RSSFeed {
	net := &http.Client{Timeout: time.Second * 3}
	r, err := net.Get(Url)
	if err != nil {
		return RSSFeed{}
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return RSSFeed{}
	}
	newFeed := RSSFeed{}
	err = xml.Unmarshal(body, &newFeed)
	if err != nil {
		return RSSFeed{}
	}
	return newFeed
}

func (cfg *apiConfig) continousFetching(limit int32, waiting time.Duration) {
	toProcess := make(chan Feed, limit)
	ctx := context.Background()
	for ; ; time.Sleep(waiting) {
		feeds, err := cfg.DB.GetNextFeedsToFetch(ctx, limit)
		if err != nil {
			fmt.Println("Scraper: Error fetching feeds")
			continue
		}
		var wg sync.WaitGroup

		rss := make([]RSSFeed, 0)

		for _, dbfeed := range feeds {
			feed := databaseFeedtoFeed(dbfeed)
			cfg.DB.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{ID: feed.ID, UpdatedAt: time.Now()})
			toProcess <- feed
		}
		for range feeds {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fd := <-toProcess
				newRSS := fetchData(fd.Url)
				newRSS.FeedID = fd.ID
				rss = append(rss, newRSS)
			}()
		}
		wg.Wait()
		cfg.processRSS(rss)
	}
}

func (cfg *apiConfig) processRSS(feeds []RSSFeed) {
	ctx := context.Background()
	for _, feed := range feeds {
		for _, item := range feed.Channel.Item {
			pubtime, err := parseTime(item.PubDate)
			if err != nil {
				fmt.Println("Error parsing date :" + item.PubDate + " error: " + err.Error())
				continue
			}
			err = cfg.DB.CreatePost(ctx, database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       sql.NullString{String: item.Title, Valid: item.Title != ""},
				Url:         item.Link,
				Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
				PublishedAt: sql.NullTime{Time: pubtime.UTC(), Valid: err == nil},
				FeedID:      feed.FeedID,
			})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	// fmt.Println("Done processing")
}

func parseTime(timestr string) (time.Time, error) {
	layouts := []string{
		"Mon, 02 Jan 2006 15:04:05 Z0700",
	}
	err := fmt.Errorf("couldn't parse")
	for _, layout := range layouts {
		outTime, err := time.Parse(layout, timestr)
		if err == nil {
			return outTime, err
		}
	}
	return time.Time{}, err
}
