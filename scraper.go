package main

import (
	"encoding/xml"
	"net/http"
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
