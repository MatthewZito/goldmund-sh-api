#!/bin/bash

# super dirty docker prune

containers=($(docker ps -a | awk '{ print $1 }'))

for container in "${containers[@]}"; do
  docker rm "$container" 2>/dev/null
done

docker image prune 2>/dev/null
docker volume prune 2>/dev/null
docker network prune 2>/dev/null