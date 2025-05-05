package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/coderjcronin/blog/internal/config"
	"github.com/coderjcronin/blog/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("Error opening database connection: %s", err)
	}
	dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		log.Fatalln("Usage: cli <command> [args...]")
		return
	}

	commandName := cliArgs[1]
	args := []string{}
	if len(cliArgs) > 2 {
		args = cliArgs[2:]
	}

	command, exists := getCommands()[commandName]
	if exists {
		err := command.callback(programState, args...)
		if err != nil {
			log.Fatalf("Error executing command: %s", err)
		}
	} else {
		log.Fatalln("Unknown command.")
	}

}
