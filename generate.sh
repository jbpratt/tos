protoc protofiles/mookies.proto --go_out=plugins=grpc:. && protoc-go-inject-tag -input=./protofiles/mookies.pb.go


python3 -m grpc_tools.protoc -I ../protofiles/ --python_out=. --grpc_python_out=. ../protofiles/mookies.proto 