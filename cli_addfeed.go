package main

import (
	"context"
	"fmt"

	"github.com/coderjcronin/blog/internal/database"
)

func commandAddFeed(s *state, user database.User, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("requires 2 arguements, <name> <url>")
	}

	returnData, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   args[0],
		Url:    args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: returnData.UserID,
		FeedID: returnData.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Created new feed '%s' with URL '%s'\nYour are now following it.\n", returnData.Name, returnData.Url)

	return nil
}
