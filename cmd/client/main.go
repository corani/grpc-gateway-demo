package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	pb "github.wdf.sap.corp/I331555/grpc-gateway-test/gen/go/helloworld"
)

func main() {
	var (
		endpoint string
		name     string
	)

	flag.StringVar(&endpoint, "endpoint", "localhost:8989", "gRPC endpoint of the greeter service")
	flag.StringVar(&name, "name", "world", "name to be greeted")
	flag.Parse()

	c, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	client := pb.NewGreeterClient(c)

	r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Greeting: %v", r.Message)
}
