protoc protofiles/mookies.proto --go_out=plugins=grpc:. && protoc-go-inject-tag -input=./protofiles/mookies.pb.go
