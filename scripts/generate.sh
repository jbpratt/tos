#!/bin/bash

set -e
pushd "$(/bin/pwd)" > /dev/null

PROTOFILE="./schema/tos.proto"
OUT_DIR="pkg/pb/"

protoc "${PROTOFILE}" \
 # --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
 # --js_out="import_style=commonjs,binary:protofiles" \
 # --ts_out="service=grpc-web:protofiles"
  --go_out "plugins=grpc:${OUT_DIR}"
#  --swift_out="$OUT_DIR" --swift_opt Visibility=Public
# --plugin=/Users/jbpratt/grpc-swift/protoc-gen-grpc-swift\
# --grpc-swift_opt=Visibilit=Public --grpc-swift_out=${OUT_DIR}

protoc-go-inject-tag -input="./pkg/pb/tos.pb.go"

popd > /dev/null
