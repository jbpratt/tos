#!/usr/bin/env bash

USER='pi'

echo Stopping system
ssh front 'killall front'
ssh kitchen 'killall kitchen'
echo Stopped
#ssh pi@mfront << EOF
#  pid=$(cat /var/tmp/mookies-pids/front.pid)
#  echo "Killing PID ${pid}"
#  pkill -9 -P ${pid}
#  sleep 1s
#  echo "Removing PID file."
#  rm -f /var/tmp/mookies-pids/*
#EOF

#ssh pi@mkitchen << EOF
#  pid=$(cat /var/tmp/mookies-pids/kitchen.pid)
#  echo "Killing PID ${pid}"
#  pkill -9 -P ${pid}
#  sleep 1s
#  echo "Removing PID file."
#  rm -f /var/tmp/mookies-pids/*
#EOF
