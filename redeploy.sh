#!/bin/sh

export MIGRATE='1'

export PGHOST='postgres'
export PGPORT='5432'
export PGNAME='music_room'
export PGUSER='alex'
export PGPASS='12345'

export RDSHOST='rds'
export RDSPORT='6379'

#git pull origin server


# Delete
docker-compose stop
docker-compose rm -f
#docker network rm mr_network

# Create
#docker network create my_app
docker-compose up -d
docker-compose logs