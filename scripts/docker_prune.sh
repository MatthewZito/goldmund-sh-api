#!/bin/bash

# super dirty docker prune

containers=($(docker ps -a | awk '{ print $1 }'))
cmds=("image" "volume" "network")

for container in "${containers[@]}"; do
  docker rm "$container" 2>/dev/null
done

for cmd in "${cmds[@]}"; do
  docker "$cmd" prune 2>/dev/null
done