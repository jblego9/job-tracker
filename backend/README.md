# Running for development
Run from the 'cmd/' directory to have environment variables setup, allowing the backend to work with databases in the 'cmd/data/' directory.

Example: `.../cmd go run ./server`

# Setup for development
Place '.env' in the 'cmd/' directory and set the environment variables.

Environment Variables:
- 'DB_PATH' to set the directory where the sqlite databases are located.
- 'PORT' to set the port that the server will listen on.
- 'NO_RESPONSE_CHECK_AFTER_N_MONTHS' to set how many months until setting an application from awaiting response to no response. Check is ran once on server startup.