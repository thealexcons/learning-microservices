package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/thealexcons/learning-microservices/01-basics/hello/hello-pb"
	"google.golang.org/grpc"
	"log"
	"os"
)

func getUserInput(s *bufio.Scanner) string {
	fmt.Println("What's your name?")
	for s.Scan() {
		return s.Text()
	}
	return ""
}


func main() {
	fmt.Println("Welcome to my (client) application!")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error occured: %v\n", err)
	}
	defer conn.Close()

	client := hellopb.NewHelloServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		inp := getUserInput(scanner)
		if inp == "" || inp == "q" {
			break
		}

		request := &hellopb.HelloRequest{Name: inp}

		resp, err := client.Hello(context.Background(), request)
		if err != nil {
			log.Fatalf("Error occured: %s\n", err)
		}
		fmt.Printf("received response from server -> '%v'\n", resp.Greeting)
	}
}
