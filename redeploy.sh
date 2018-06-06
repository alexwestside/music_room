#!/bin/sh

PGHOST='postgres'
PGPORT='5432'
PGNAME='music-room'
PGUSER='postgres'
PGPASS='12345'
MIGRATE='1'

git pull origin server

docker-compose stop
docker-compose rm -f
docker-compose up -d
