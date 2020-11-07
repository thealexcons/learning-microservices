package main

import (
	"github.com/thealexcons/learning-microservices/01-basics/hello/hello-pb"
	"github.com/thealexcons/learning-microservices/01-basics/hello/hello-server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error occured: %v\n", err)
	}
	log.Println("gRPC server is now listening...")

	gRPCServer := grpc.NewServer()
	s := server.NewServer()	// our service
	hellopb.RegisterHelloServiceServer(gRPCServer, s)

	gRPCServer.Serve(lis)
}
