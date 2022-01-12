package gateway

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.wdf.sap.corp/I331555/grpc-gateway-test/gen/go/helloworld"
	"github.wdf.sap.corp/I331555/grpc-gateway-test/internal/server"
)

func RegisterGateway(r *mux.Router, s server.Server, prefix string) {
	c, err := grpc.DialContext(
		context.Background(), fmt.Sprintf("0.0.0.0:%d", s.Port()),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	handler := runtime.NewServeMux()

	if err := pb.RegisterGreeterHandler(context.Background(), handler, c); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	r.NewRoute().PathPrefix(prefix).Handler(handler)
}
