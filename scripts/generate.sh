#!/bin/bash

set -e
pushd "$(/bin/pwd)" > /dev/null

PROTOFILE="./protofiles/tos.proto"
OUT_DIR="."

protoc "${PROTOFILE}" \
  --go_out "plugins=grpc:${OUT_DIR}" \
  --js_out "import_style=commonjs:${OUT_DIR}" \
  --grpc-web_out "import_style=commonjs,mode=grpcwebtext:${OUT_DIR}"
  --swift_out="$OUT_DIR" --swift_opt Visibility=Public

protoc-go-inject-tag -input="./protofiles/tos.pb.go"

popd > /dev/null
