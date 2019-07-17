package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/tahsinrahman/grpc-demo/helloworld"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "hello " + req.Name}, nil
}

func (s *server) StreamHello(req *helloworld.HelloRequest, stream helloworld.HelloService_StreamHelloServer) error {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		err := stream.Send(&helloworld.HelloReply{
			Message: "hello " + req.Name,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	helloworld.RegisterHelloServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
