#!/bin/bash

# These are the environmental variables you should set before running allwrite.
# You can set these in something like Upstart or in your current shell by doing:

# source ./creds.sh
# ------------------------------------------------------------------------------

# Basically, you just need to make sure these get set somehow, somewhere.

# The ID of the base directory for the docs (you can grab it from the URL in
# Drive).
export ACTIVE_DIR="xxxxxxxxxxxxxxxxxxx"

# Path to your Google client secret json file.
export CLIENT_SECRET="$PWD/client_secret.json"

# The storage system to use - currently postgres is the only option.
export STORAGE="postgres"
export PG_USER="root"
export PG_PASS="root"
export PG_DB="allwrite"
export PG_HOST="localhost"

# Specify the port to run the application on.
export PORT=":8000"

# Only needed if listening on port 443. Used for certbot.
export DOMAIN="my-domain.com"
export CERTBOT_EMAIL="engineering@your-company.com"

# How often Google is queried for updates specified in milliseconds.
export FREQUENCY="300000"
