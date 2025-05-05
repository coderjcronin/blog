package main

import (
	"context"
	"fmt"
)

func commandReset(s *state, args ...string) error {
	err := s.db.DelAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("All users deleted.")
	return nil
}
