package main

type cliCommand struct {
	name        string
	description string
	callback    func(*state, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"register": {
			name:        "register",
			description: "Register new user in database",
			callback:    commandRegister,
		},
		"login": {
			name:        "login",
			description: "Sets the current user in config",
			callback:    commandLogin,
		},
		"users": {
			name:        "users",
			description: "Lists all registered users, indicates current user",
			callback:    commandUsers,
		},
		"agg": {
			name:        "agg",
			description: "Aggregates RSS feeds followed based on <time_between_reqs> (duration string)",
			callback:    middlewareLoggedIn(commandAgg),
		},
		"addfeed": {
			name:        "addfeed",
			description: "Adds feed with <name> and <url>",
			callback:    middlewareLoggedIn(commandAddFeed),
		},
		"feeds": {
			name:        "feeds",
			description: "List feeds and feed creators",
			callback:    commandFeeds,
		},
		"follow": {
			name:        "follow",
			description: "Follow a feed with <url>",
			callback:    middlewareLoggedIn(commandFollow),
		},
		"following": {
			name:        "following",
			description: "List the feeds the current user is following",
			callback:    middlewareLoggedIn(commandFollowing),
		},
		"unfollow": {
			name:        "unfollow",
			description: "Unfollow feed by <url>",
			callback:    middlewareLoggedIn(commandUnfollow),
		},
		"dburl": {
			name:        "dburl",
			description: "Sets the current DB url in config",
			callback:    commandDbUrl,
		},
		"reset": {
			name:        "reset",
			description: "Reset all users in database (for testing)",
			callback:    commandReset,
		},
	}
}
