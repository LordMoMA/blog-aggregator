package main

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lordmoma/blog-aggregator/internal/database"
)

type RssFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Item        []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeedData(feedURL string) (*RssFeed, error) {
	resp, err := http.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss RssFeed
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}

func fetchFeedsWorker(db *database.Queries, concurrency int32) {
	// Set up ticker for interval
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	// Set up wait group for processing feeds concurrently
	var wg sync.WaitGroup

	for range ticker.C {
		// Get the next n feeds to fetch from the database
		ctx := context.Background()
		feeds, err := db.GetNextFeedsToFetch(ctx, concurrency)
		if err != nil {
			log.Printf("Error getting next feeds to fetch: %v\n", err)
			continue
		}

		if len(feeds) == 0 {
			log.Printf("No feeds to fetch")
			continue
		}

		log.Printf("Fetching %d feeds...\n", len(feeds))

		// Fetch and process all the feeds at the same time
		wg.Add(len(feeds))
		for _, feed := range feeds {
			go func(feed database.Feed) {
				defer wg.Done()

				_, err := db.MarkFeedFetched(ctx, feed.ID)
				if err != nil {
					log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
					return
				}

				rss, err := FetchFeedData(feed.Url)
				if err != nil {
					log.Printf("Error fetching feed %s: %v\n", feed.Url, err)
					return
				}

				log.Printf("Feed %s:\n", rss.Channel.Title)

				for _, item := range rss.Channel.Item {

					dateStr := item.PubDate
					layout := "Mon, 02 Jan 2006 15:04:05 -0700"

					date, err := time.Parse(layout, dateStr)
					if err != nil {
						log.Printf("Error parsing time: %v\n", err)
						return
					}

					post := database.CreatePostParams{
						ID:          uuid.New(),
						FeedID:      feed.ID,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Title:       item.Title,
						Description: item.Description,
						Url:         item.Link,
						PublishedAt: date,
					}

					p, err := db.CreatePost(ctx, post)
					if err != nil {
						log.Printf("Error creating post: %v\n", err)
						return
					}
					log.Printf("- %s\n", p)
				}
			}(feed)
		}
		// Wait for all feeds to finish processing
		wg.Wait()
	}
}
