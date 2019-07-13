#!/bin/sh
while true
do
  go run cli/main.go -a :50051 order
  echo "order placed"
sleep $((RANDOM % 350))
done
