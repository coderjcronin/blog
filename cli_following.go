package main

import (
	"context"
	"fmt"

	"github.com/coderjcronin/blog/internal/database"
)

func commandFollowing(s *state, user database.User, args ...string) error {

	feeds, err := s.db.GetFeedsFollowing(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feeds) < 1 {
		fmt.Println("You are not following any feeds. Add some!")
	} else {
		fmt.Printf("You are following:\n")

		for _, feed := range feeds {
			fmt.Printf("\t* %s\n", feed.FeedName)
		}
	}

	return nil
}
