package main

import (
	"context"

	"github.com/coderjcronin/blog/internal/database"
)

func middlewareLoggedIn(handler func(s *state, user database.User, args ...string) error) func(*state, ...string) error {
	return func(s *state, args ...string) error {
		// Get the current user - this is the code you're avoiding duplicating
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		// Call the original handler with the user
		return handler(s, user, args...)
	}
}
