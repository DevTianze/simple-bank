#!/bin/sh
# Start the service
set -e

echo "run db migration"
echo "ls /app"
ls -l /app/migration
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"