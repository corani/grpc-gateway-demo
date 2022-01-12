package client

import (
	"context"
	"log"

	pb "github.wdf.sap.corp/I331555/grpc-gateway-test/gen/go/helloworld"
	"google.golang.org/grpc"
)

type Client interface {
	SayHello(string) string
}

type client struct {
	client pb.GreeterClient
}

func NewClient(endpoint string) Client {
	c, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	return &client{
		client: pb.NewGreeterClient(c),
	}
}

func (c *client) SayHello(name string) string {
	r, err := c.client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	return r.Message
}
