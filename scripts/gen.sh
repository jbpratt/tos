#!/bin/sh

#PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
OUT_DIR="./generated"
protoc -I=. mookies.proto \
--js_out=import_style=commonjs:$OUT_DIR \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT_DIR

#docker build -t tos/envoy -f ./envoy.Dockerfile .
