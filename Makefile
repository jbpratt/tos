PKG := "github.com/jbpratt78/tos"
GOPATH = /home/jbpratt/go
SERVER_OUT := "bin/server"
SERVER_PKG_BUILD := "${PKG}"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

dep:
	@go get -u

lint:
	@golint -set_exit_status ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

deploy: start
	$(sh ./scripts/start_system.sh)

server: gen
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

clean:
	@rm $(SERVER_OUT) $(FRONT_CLIENT_OUT) $(BACK_CLIENT_OUT)

test:
	@go test -short ${PKG_LIST}

start: server
	./bin/server

gen:
	@protoc -I/usr/local/include -I. \
			-I${GOPATH}/src \
			-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:. \
		protofiles/mookies.proto
	@protoc-go-inject-tag -input=protofiles/mookies.pb.go
	@protoc -I/usr/local/include -I. \
			-I${GOPATH}/src \ 
			-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--grpc-gateway_out=logtostderr=true:. protofiles/mookies.proto
	@mockgen -source=protofiles/mookies.pb.go > mock/proto_mock.go
