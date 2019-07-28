#!/usr/bin/env bash

echo "Starting server..."
docker-compose up -d
ssh server << EOF
  cd go/src/github.com/jbpratt78/tos
  docker-compose up -d
  sleep 2s
EOF
echo "Server started..."

echo "Starting front..."
DISPLAY=:0 go/bin/front > ~/front.log 2>&1 &
echo $! > /var/tmp/mookies-pids/front.pid
sleep 2s
echo "Front started..."

ssh kitchen << EOF
  DISPLAY=:0 kitchen > ~/kitchen.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/kitchen.pid
  sleep 2s
EOF
