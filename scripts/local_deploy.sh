#!/usr/bin/env bash

make test
docker-compose up -d

### Not in use (Docker)
#./bin/server --database=server/mookies.db > ~/server.log 2>&1 &
#echo $! > /var/tmp/mookies-pids/server.pid

ssh front <<EOF
  GO111MODULE=on /usr/local/go/bin/go install github.com/jbpratt/tos/front
  DISPLAY=:0 go/bin/front > ~/front.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/front.pid
  sleep 2s
EOF

ssh kitchen <<EOF
  GO111MODULE=on /usr/local/go/bin/go install github.com/jbpratt/tos/kitchen
  DISPLAY=:0 kitchen > ~/kitchen.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/kitchen.pid
  sleep 2s
EOF
