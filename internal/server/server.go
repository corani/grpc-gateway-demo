package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.wdf.sap.corp/I331555/grpc-gateway-test/gen/go/helloworld"
	"google.golang.org/grpc"
)

type Server interface {
	Port() int
	Start()
}

type server struct {
	pb.UnimplementedGreeterServer

	port int
}

func NewServer(port int) Server {
	return &server{port: port}
}

func (s *server) Port() int {
	return s.port
}

func (s *server) Start() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	gs := grpc.NewServer()

	pb.RegisterGreeterServer(gs, s)

	log.Printf("Serving gRPC on 0.0.0.0:%d", s.port)
	go func() {
		log.Fatal(gs.Serve(l))
	}()
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("[grpc] SayHello to %q", in.Name)

	return &pb.HelloReply{Message: "hello, " + in.Name}, nil
}
