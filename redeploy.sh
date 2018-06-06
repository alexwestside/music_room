#!/bin/sh

MIGRATE='1'

PGHOST='postgres'
PGPORT='5432'
PGNAME='music_room'
PGUSER='alex'
PGPASS='12345'

RDSHOST='rds'
RDSPORT='6379'

git pull origin server

docker-compose stop
docker-compose rm -f
docker-compose up -d -e PGHOST -e PGPORT -e PGNAME -e PGUSER -e PGPASS -e RDSHOST -e RDSPORT -e MIGRATE
