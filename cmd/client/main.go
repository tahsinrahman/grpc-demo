package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tahsinrahman/grpc-demo/helloworld"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := helloworld.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "tahsin"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(reply.Message)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := c.StreamHello(ctx, &helloworld.HelloRequest{Name: "tahsin"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		fmt.Println(msg.Message)
	}
}
