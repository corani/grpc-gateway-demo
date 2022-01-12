package main

import (
	"flag"
	"log"

	"github.wdf.sap.corp/I331555/grpc-gateway-test/internal/client"
)

func main() {
	var (
		endpoint string
		name     string
	)

	flag.StringVar(&endpoint, "endpoint", "localhost:8989", "gRPC endpoint of the greeter service")
	flag.StringVar(&name, "name", "world", "name to be greeted")
	flag.Parse()

	c := client.NewClient(endpoint)

	log.Printf("Greeting: %v", c.SayHello(name))
}
