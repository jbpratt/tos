#!/usr/bin/env bash

PKG='github.com/jbpratt78/tos'
OUT_DIR='bin/'
USER='pi'
FRONT_CONFIG=()
FRONT_CONFIG+=('--addr="192.168.1.104:50051"')

KITCHEN_CONFIG=()

make test
docker-compose up -d

### Not in use (Docker)
#./bin/server --database=server/mookies.db > ~/server.log 2>&1 &
#echo $! > /var/tmp/mookies-pids/server.pid


ssh -t ${USER}@mfront << EOF
  GO111MODULE=on /usr/local/go/bin/go install github.com/jbpratt78/tos/front
  DISPLAY=:0 go/bin/front --addr="192.168.1.104:50051" > ~/front.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/front.pid
EOF
