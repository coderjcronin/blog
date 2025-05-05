package main

import (
	"context"
	"fmt"
)

func commandRegister(programState *state, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid number of arguments")
	}

	newUserName := args[0]

	_, err := programState.db.CreateUser(context.Background(), newUserName)
	if err != nil {
		return err
	}

	err = commandLogin(programState, newUserName)
	if err != nil {
		return err
	}

	return nil
}
