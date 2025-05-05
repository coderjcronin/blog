package main

import "fmt"

func commandHelp(progamState *state, args ...string) error {
	commands := getCommands()
	fmt.Println("-- HELP --")

	for _, command := range commands {
		fmt.Printf("%s\t\t%s\n", command.name, command.description)
	}

	return nil
}
