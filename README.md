# gator
Gator is an RSS aggregator.

To use gator, you'll need: Postgres and Go.
macOS with brew:

brew install postgresql@15

Linux / WSL (Debian):

sudo apt update
sudo apt install postgresql postgresql-contrib
Next, install the gator CLI:
go install gator

Gator is only configured to run locally. It needs a config file on your computer to work.
If there isn't one already, create a ".gatorconfig.json" file in your home directory, like so:

{
    "db_url": "connection_string_goes_here",
    "current_user_name": "username_goes_here"
}
This keeps track of who is currently logged in, and the connection credentials for the Postgres database.

A connection string looks like this: protocol://username:password@host:port/database


Gator Commands:

addfeed <name> <url> - add a feed to the database's list of known feeds.

agg <time duration string: 1h, 1m, 1s etc> - begin aggregating a users feeds for browsing later. CAUTION - do not set the refresh duration to be too short

browse <number of most recent posts> - browse posts from a user's feeds

feeds - retrieves a list of feeds on the server

follow <url> - follow a feed at a given URL

following - retrieves a list of the current user's followed feeds.

login <name> - log in as the provided user.

register <name> - register a new user.

reset - resets the database - probably don't do this

unfollow - unfollow a feed at the given URL.

users - retrieves a list of registered users on the database.