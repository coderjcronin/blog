 # Gator Blog Aggregator
 ## boots.dev guided project

 ### Requirements
 Go and postgres

 ### Installation
 Enter the following command into your command line: `go install github.com/coderjcronin/blog`

 ### Config File
 Create the following file in your home directory (`/~`): `.gatorconfig.json` with the following information:
 `{"db_url":"postgres://username:password@host:port/database?sslmode=disable","current_user_name":""}`
 Replace username, password, host, port, and database with the appropriate information. Leave `current_user_name` blank.

 ### Commands
 Run `./blog help` to see a list of commands.
 `./blog register <username>` to register a new user and set them as active
 `./blog login <username>` to set an already registered user as active
 `./blog addfeed <title> <url>` to register and follow a new feed
 `./blog follow <url>` to follow a feed added already by another user
 `./blog agg <duration_string>` to aggregate posts from feeds for current active user to database; please note too short a duration will cause inappropriate constant requests and will get your IP banned by the host. You've been warned, FAFO.
 `./blog browse <optional_limit>` to browse stored posts for the active user's feeds, not specifying limit will return 2 latest posts.