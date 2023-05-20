package pb

//go:generate cd konfig-go
//go:generate protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --proto_path=./pb ./pb/konfig.proto
