#!/usr/bin/env bash

USER='pi'

make test
docker-compose up -d

### Not in use (Docker)
#./bin/server --database=server/mookies.db > ~/server.log 2>&1 &
#echo $! > /var/tmp/mookies-pids/server.pid


ssh -t ${USER}@front << EOF
  GO111MODULE=on /usr/local/go/bin/go install github.com/jbpratt78/tos/front
  DISPLAY=:0 go/bin/front > ~/front.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/front.pid
EOF

ssh -t ${USER}@kitchen << EOF
  GO111MODULE=on /usr/local/go/bin/go install github.com/jbpratt78/tos/kitchen
  DISPLAY=:0 go/bin/kitchen > ~/kitchen.log 2>&1 &
  echo $! > /var/tmp/mookies-pids/kitchen.pid
EOF
