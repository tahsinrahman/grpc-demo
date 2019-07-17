package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tahsinrahman/grpc-demo/helloworld"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := helloworld.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, "localhost:5050", opts)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
