package main

import (
	"context"
	"fmt"

	"github.com/coderjcronin/blog/internal/database"
)

func commandUnfollow(s *state, user database.User, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing url arguement")
	}

	err := s.db.DeleteFollowByUrl(context.Background(), database.DeleteFollowByUrlParams{
		Url:    args[0],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Unfollowed.")

	return nil
}
