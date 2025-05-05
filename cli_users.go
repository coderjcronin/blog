package main

import (
	"context"
	"fmt"
)

func commandUsers(s *state, args ...string) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) < 1 {
		return fmt.Errorf("no users")
	}

	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf(" * %s (current)\n", user)
		} else {
			fmt.Printf(" * %s\n", user)
		}
	}

	return nil
}
