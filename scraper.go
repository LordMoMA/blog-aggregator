package main

import (
	"context"
	"database/sql"
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

					publishedAt := sql.NullTime{}
					if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
						publishedAt = sql.NullTime{
							Time:  t,
							Valid: true,
						}
					}

					post := database.CreatePostParams{
						ID:        uuid.New(),
						FeedID:    feed.ID,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Title:     item.Title,
						Description: sql.NullString{
							String: item.Description,
							Valid:  true,
						},
						Url:         item.Link,
						PublishedAt: publishedAt,
					}

					_, err := db.CreatePost(ctx, post)
					if err != nil {
						log.Printf("Error creating post: %v\n", err)
						continue
					}
					log.Printf("Feed %s collected, %v posts found", feed.Name, len(rss.Channel.Item))
				}
			}(feed)
		}
		// Wait for all feeds to finish processing
		wg.Wait()
	}
}
