PKG := "github.com/jbpratt78/tos"
SERVER_OUT := "bin/server"
SERVER_PKG_BUILD := "${PKG}"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

server: dep gen
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

dep:
	@go get github.com/golang/protobuf/protoc-gen-go
	@go get ./...

lint:
	@staticcheck ./...

deploy: start
	$(sh ./scripts/start_system.sh)

clean:
	@rm $(SERVER_OUT) $(FRONT_CLIENT_OUT) $(BACK_CLIENT_OUT)

test:
	@go test -short ${PKG_LIST}

start:
	./bin/server

gen:
	@protoc --go_out=plugins=grpc:. protofiles/tos.proto
	@protoc-go-inject-tag -input=protofiles/tos.pb.go > /dev/null 2>&1
