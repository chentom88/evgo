package main

import (
	"api/api_proto"
	"api/apiservice"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	api_proto.RegisterDestinationAPIServer(grpcServer, apiservice.NewDestinationService())
	go grpcServer.Serve(listener)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = api_proto.RegisterDestinationAPIHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server")
	http.ListenAndServe(":8080", mux)
}
