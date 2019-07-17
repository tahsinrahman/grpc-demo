### generate rpc code
protoc -I ./helloworld/ --go_out=plugins=grpc:./helloworld/ ./helloworld/helloworld.proto

### run the server
go run cmd/server/main.go

### run the client
go run cmd/client/main.go