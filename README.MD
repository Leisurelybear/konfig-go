1. install protoc tools
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
go get google.golang.org/genproto/...
```

2. copy google api to pb directory
3. generate pb.go
``` shell
cd konfig-go
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --proto_path=./pb ./pb/konfig.proto
```