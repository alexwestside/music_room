#!/bin/sh

MIGRATE='1'

PGHOST='postgres'
PGPORT='5432'
PGNAME='music-room'
PGUSER='postgres'
PGPASS='12345'

RDSHOST='rds'
RDSPORT='6379'

git pull origin server

docker-compose stop
docker-compose rm -f
docker-compose up -d
