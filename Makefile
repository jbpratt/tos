PKG := "github.com/jbpratt78/mookies-tos"
SERVER_OUT := "bin/server"
FRONT_CLIENT_OUT := "bin/front"
BACK_CLIENT_OUT := "bin/kitchen"
SERVER_PKG_BUILD := "${PKG}/server"
FRONT_CLIENT_PKG_BUILD := "${PKG}/front"
BACK_CLIENT_PKG_BUILD := "${PKG}/kitchen"

all: server front back

dep: ## Get the dependencies
	@go get -u

server: dep protofiles/mookies.pb.go
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

front: dep protofiles/mookies.pb.go
	@go build -i -v -o $(FRONT_CLIENT_OUT) $(FRONT_CLIENT_PKG_BUILD)

back: dep protofiles/mookies.pb.go
	@go build -i -v -o $(BACK_CLIENT_OUT) $(BACK_CLIENT_PKG_BUILD)

clean:
	@rm $(SERVER_OUT) $(FRONT_CLIENT_OUT) $(BACK_CLIENT_OUT)

start: back 
	@docker-compose up -d
	rm -rf ./bin/server

proto:
	protoc protofiles/mookies.proto --go_out=plugins=grpc:.
	protoc-go-inject-tag -input=protofiles/mookies.pb.go
	python3 -m grpc_tools.protoc -I protofiles/ --python_out=printing/ \
		--grpc_python_out=printing/ protofiles/mookies.proto
