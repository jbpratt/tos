#!/bin/bash

set -e
pushd "$(/bin/pwd)" > /dev/null

PROTOFILE="./protofiles/tos.proto"
OUT_DIR="./protofiles"

protoc "${PROTOFILE}" \
  --go_out "plugins=grpc:${OUT_DIR}" \
  --js_out "import_style=commonjs:${OUT_DIR}" \
  --grpc-web_out "import_style=commonjs,mode=grpcwebtext:${OUT_DIR}"

protoc-go-inject-tag -input="${PROTOFILE}"

popd > /dev/null
