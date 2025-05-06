package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/coderjcronin/blog/internal/database"
)

func commandBrowse(s *state, user database.User, args ...string) error {
	var limiter int32
	if len(args) < 1 {
		limiter = 2
	} else {
		prepInt, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		limiter = int32(prepInt)
	}

	posts, err := s.db.GetUserPosts(context.Background(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  limiter,
	})
	if err != nil {
		return err
	}

	fmt.Println("Posts:")
	for _, post := range posts {
		fmt.Printf("%s\t%s\n%s\n%s\n\n", post.Title, post.PublishedAt.Time.String(), post.Url, post.Description.String)
	}

	return nil
}
