package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/araddon/dateparse"
	"github.com/coderjcronin/blog/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	for i, item := range rss.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rss.Channel.Item[i] = item
	}

	return rss, nil
}

func addPost(s *state, item RSSItem, feed uuid.UUID) error {

	description := sql.NullString{
		String: item.Description,
		Valid:  true,
	}

	t, err := dateparse.ParseAny(item.PubDate)
	if err != nil {
		t = time.Now()
	}
	timeStamp := sql.NullTime{
		Time:  t,
		Valid: true,
	}

	returnData, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
		Title:       item.Title,
		Url:         item.Link,
		Description: description,
		PublishedAt: timeStamp,
		FeedID:      feed,
	})
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			// Handle the unique constraint violation
			// For example, you can return a specific error message to the client
			fmt.Printf("%s already exists, skipping...\n", item.Title)
			return nil
		} else {
			return err
		}
	}

	fmt.Printf("Added '%s' to posts.\n", returnData.Title)

	return nil
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
		return nil
	} else {
		fmt.Printf("Channel %s:\n", rss.Channel.Title)
		for _, item := range rss.Channel.Item {
			err = addPost(s, item, nextFeed.ID)
			if err != nil {
				fmt.Println(err)
			}
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

}
