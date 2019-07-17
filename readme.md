### generate rpc code
protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ -I ./helloworld/ --go_out=plugins=grpc:./helloworld/ ./helloworld/helloworld.proto

protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ -I ./helloworld/ --grpc-gateway_out=logtostderr=true:./helloworld/ ./helloworld/helloworld.proto

protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ -I ./helloworld/ --swagger_out=logtostderr=true:./helloworld/ ./helloworld/helloworld.proto

### run the server
go run cmd/server/main.go

### run the client
go run cmd/client/main.go

### run the proxy
go run cmd/proxy/main.go

```bash
$ curl -XPOST localhost:8081/hello -d '{ "name" : "tahsin"}'
{"message":"hello tahsin"}

$ curl -XGET localhost:8081/hello/tahsin
{"message":"hello tahsin"}


$ curl -XGET localhost:8081/sayhello/tahsin
{"message":"hello tahsin"}

```
