#!/bin/bash

PROTOCOL='http'
BASE_URI='localhost'
PORT=5000

URL="$PROTOCOL://$BASE_URI:$PORT/"

[[ ! -z "$1" ]] && URL="$URL$1"

curl -D - "$URL"
echo ""