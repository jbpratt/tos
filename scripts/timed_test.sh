#!/bin/sh
i=0
while true; do
	i=$((i + 1))
	go run cli/main.go -a :50051 order
	echo "order placed $i"
	sleep $((RANDOM % 1200))
done
