PKG := "./cmd/server"
SERVER_OUT := "bin/server"
SERVER_PKG_BUILD := "${PKG}"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
PB_OUT := "pkg/pb"

server: dep
	@go build -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

dep:
	@go get github.com/golang/protobuf/protoc-gen-go
	@go get ./...

lint:
	@staticcheck ./...

clean:
	@rm $(SERVER_OUT) $(FRONT_CLIENT_OUT) $(BACK_CLIENT_OUT)

test:
	@go test -short ${PKG_LIST}

start:
	./bin/server

gen:
	@protoc --go_out=plugins=grpc:. schema/tos.proto
	@protoc-go-inject-tag -input=protofiles/tos.pb.go > /dev/null 2>&1
	@mv protofiles/tos.pb.go $(PB_OUT)
