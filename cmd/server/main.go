package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.wdf.sap.corp/I331555/grpc-gateway-test/gen/go/helloworld"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "hello, " + in.Name}, nil
}

func startGRPC(port int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	helloworld.RegisterGreeterServer(s, &server{})

	log.Printf("Serving gRPC on 0.0.0.0:%d", port)
	go func() {
		log.Fatal(s.Serve(l))
	}()
}

func RegisterGateway(r *mux.Router, port int, prefix string) {
	c, err := grpc.DialContext(
		context.Background(), fmt.Sprintf("0.0.0.0:%d", port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	handler := runtime.NewServeMux()

	if err := helloworld.RegisterGreeterHandler(context.Background(), handler, c); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	r.NewRoute().PathPrefix(prefix).Handler(handler)
}

func RegisterOpenAPIHandler(r *mux.Router, path string) {
	spec, err := openapi3.NewLoader().LoadFromFile("docs/api/helloworld/hello_world.swagger.json")
	if err != nil {
		log.Fatalf("failed to load: %v", err)
	}

	r.NewRoute().Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(spec); err != nil {
			log.Printf("failed to encode: %v", err)
		}
	})
}

func RegisterStaticContent(r *mux.Router, prefix string) {
	r.NewRoute().PathPrefix(prefix).Handler(
		http.StripPrefix(prefix,
			http.FileServer(http.Dir("static/"))))
}

func main() {
	startGRPC(8989)

	r := mux.NewRouter()
	RegisterOpenAPIHandler(r, "/openapi3.json")
	RegisterGateway(r, 8989, "/v1")
	RegisterStaticContent(r, "/static/")

	srv := &http.Server{
		Addr:    ":8990",
		Handler: r,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8990")
	log.Fatalln(srv.ListenAndServe())
}
