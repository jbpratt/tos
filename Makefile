PKG := "github.com/jbpratt78/tos"
GOPATH = /home/jbpratt/go
SERVER_OUT := "bin/server"
FRONT_CLIENT_OUT := "bin/front"
BACK_CLIENT_OUT := "bin/kitchen"
SERVER_PKG_BUILD := "${PKG}"
FRONT_CLIENT_PKG_BUILD := "${PKG}/front"
BACK_CLIENT_PKG_BUILD := "${PKG}/kitchen"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

all: server front back

dep:
	@go get -u

lint:
	@golint -set_exit_status ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

deploy: start
	$(sh ./scripts/start_system.sh)

server: dep protofiles/mookies.pb.go
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

front: dep protofiles/mookies.pb.go
	@go build -i -v -o $(FRONT_CLIENT_OUT) $(FRONT_CLIENT_PKG_BUILD)

back: dep protofiles/mookies.pb.go
	@go build -i -v -o $(BACK_CLIENT_OUT) $(BACK_CLIENT_PKG_BUILD)

clean:
	@rm $(SERVER_OUT) $(FRONT_CLIENT_OUT) $(BACK_CLIENT_OUT)

test:
	@go test -short ${PKG_LIST}

start:
	@docker-compose up -d
	@rm -rf ./bin/server

gen:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		protofiles/mookies.proto
	@protoc-go-inject-tag -input=protofiles/mookies.pb.go
	@protoc -I/usr/local/include -I. \
  	-I${GOPATH}/src \
  	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  	--grpc-gateway_out=logtostderr=true:. \
		protofiles/mookies.proto
	@mockgen -source=protofiles/mookies.pb.go > mock/proto_mock.go
