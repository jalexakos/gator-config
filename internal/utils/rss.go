package utils

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error building new HTTP request %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error fetching RSS feed %w", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body %w", err)
	}

	var rssData RSSFeed

	xmlUnmarshalErr := xml.Unmarshal(data, &rssData)
	if xmlUnmarshalErr != nil {
		return nil, fmt.Errorf("Error unmarshaling response body into RSS Data %w", err)
	}

	return &rssData, nil
}

// func EscapeHTMLChars()
