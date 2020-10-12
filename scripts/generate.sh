#!/bin/bash

set -e
pushd "$(/bin/pwd)" > /dev/null

PROTOFILE="./protofiles/tos.proto"
OUT_DIR="pkg/pb/"

protoc "${PROTOFILE}" \
  --go_out "plugins=grpc:${OUT_DIR}"
#  --swift_out="$OUT_DIR" --swift_opt Visibility=Public

protoc-go-inject-tag -input="./pkg/pb/tos.pb.go"

popd > /dev/null
