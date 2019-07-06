#!/bin/sh
while true
do
  go run cli/main.go -a server:50051 order
  echo "order placed"
sleep 300
done
