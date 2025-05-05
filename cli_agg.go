package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/coderjcronin/blog/internal/database"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss *RSSFeed
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, err
	}

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	for _, item := range rss.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return rss, nil
}

func scrapeFeeds(s *state, user database.User) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background(), user.ID)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()

	rss, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}

	if len(rss.Channel.Item) < 1 {
		fmt.Printf("%s has no items to view\n", rss.Channel.Title)
	} else {
		fmt.Printf("Channel %s:\n", rss.Channel.Title)
		for _, item := range rss.Channel.Item {
			fmt.Printf("\t- %s\n", item.Title)
		}
	}

	return nil
}

func commandAgg(s *state, user database.User, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid number of arguements, missing duration between scrapes")
	}

	timeBetweenRequests, err := time.ParseDuration(args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s, user)
	}

	return nil
}
