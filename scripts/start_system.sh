#!/usr/bin/env bash

docker-compose up -d

ssh front << EOF
  DISPLAY=:0 go/bin/front > ~/front.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/front.pid
  sleep 2s
EOF

ssh kitchen << EOF
  DISPLAY=:0 kitchen > ~/kitchen.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/kitchen.pid
  sleep 2s
EOF
