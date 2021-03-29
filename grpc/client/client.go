package main

import (
	"context"
	myarc "github.com/Michael-Johy/go-pkgs-practise/grpc/proto/proto/grpc"
	"google.golang.org/grpc"
	"log"
)

func main() {

	address := "127.0.0.1:8080"
	conn, e := grpc.Dial(address, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("failed to dial %v", address)
	}
	defer conn.Close()

	client := myarc.NewHelloWorldClient(conn)
	resp, e := client.SayHelloWorld(context.Background(), &myarc.HelloRequest{
		Name: "yt",
	})
	if e != nil {
		log.Fatalf("failed to get res for param %v", 1)
	}
	log.Println(resp)

}
