package main

import "fmt"

func commandDbUrl(programState *state, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid amount of arguements")
	}

	err := programState.cfg.SetDB(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Set DB URL to %s", args[0])

	return nil
}
