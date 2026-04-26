# Running
Run from the 'cmd/' directory to have environment variables setup, allowing the backend to work with databases in the 'cmd/data/' directory.

Example: `.../cmd go run ./server`

# Environment setup
Place '.env' in the 'cmd/' directory and set the environment variables.

Environment Variables:
- 'DB_PATH' to set the directory where the sqlite databases are located.
- 'PORT' to set the port that the server will listen on.