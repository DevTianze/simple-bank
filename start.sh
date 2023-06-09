#!/bin/sh
# Start the service
set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"