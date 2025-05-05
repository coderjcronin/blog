package main

import (
	"context"
	"fmt"

	"github.com/coderjcronin/blog/internal/database"
)

func commandFollow(s *state, user database.User, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid amount of arguements, need url")
	}

	feed, err := s.db.LookupFeedByUrl(context.Background(), args[0])
	if err != nil {
		return err
	}

	followResult, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Success\n%s is now following %s\n", followResult.UserName, followResult.FeedName)

	return nil
}
