## How to dev
1. install protobuf tools
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
go get google.golang.org/genproto/...
```

2. copy google api to pb directory
3. generate pb.go
``` shell
cd konfig-go # in project directory
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --proto_path=./pb ./pb/konfig.proto
```

## How to run
1. config DB, copy config, then modify it
```shell
cp conf/config.toml.template conf/config.toml
```

2. run server
```go
go run main.go
```

