package main

import (
	"context"
	"fmt"
)

func commandLogin(programState *state, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no user specified")
	}

	name := args[0]

	_, err := programState.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = programState.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("%s has been set as current db user.\n", name)

	return nil
}
