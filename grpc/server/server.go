package main

import (
	"context"
	mygrpc "github.com/Michael-Johy/go-pkgs-practise/grpc/proto/proto/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	mygrpc.UnimplementedHelloWorldServer
}

func (*server) SayHelloWorld(ctx context.Context, in *mygrpc.HelloRequest) (*mygrpc.HelloResponse, error) {
	return &mygrpc.HelloResponse{Message: in.GetName() + ", World"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	mygrpc.RegisterHelloWorldServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	if e := s.Serve(lis); e != nil {
		log.Fatalf("failed to server %v", e)
	}
}
