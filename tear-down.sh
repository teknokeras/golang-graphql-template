#!/bin/bash

docker-compose stop
docker-compose down
docker rmi webapp-go
docker volume rm $(docker volume ls -f dangling=true -q)

