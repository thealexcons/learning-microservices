package server

import (
	"context"
	"github.com/thealexcons/learning-microservices/01-basics/hello/hello-pb"
)

type server struct {
}

func NewServer() *server {
	return &server{}
}

// Implementation of the Hello RPC call
func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := req.Name
	response :=&hellopb.HelloResponse{
		Greeting: "Hello there, " + name,
	}
	return response, nil
}


