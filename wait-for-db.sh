#!/bin/sh
set -e

until pg_isready -h "db" -U "me" -d "development"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 3
done

>&2 echo "Postgres is up - executing command"
exec "$@"
