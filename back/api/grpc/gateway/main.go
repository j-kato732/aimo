package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	errdetails "github.com/j-kato732/aimo/errors"
	gw "github.com/j-kato732/aimo/proto"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

const (
	port = ":9002"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC servcer endpoint
	// NOde: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(errdetails.CustomHTTPError),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterAimoHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	newMux := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT"}),
		handlers.AllowedOrigins([]string{"http://localhost:9001"}),
		handlers.AllowedHeaders([]string{"content-type", "authorization"}),
	)(mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("server listening at %v\n", port)
	return http.ListenAndServe(port, newMux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
