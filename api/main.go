package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"git.neds.sh/matty/entain/api/proto/racing"
	"git.neds.sh/matty/entain/api/proto/sports"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	apiEndpoint        = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	grpcRacingEndpoint = flag.String("grpc-endpoint", "localhost:9000", "gRPC server endpoint")
	grpcSportsEndpoint = flag.String("grpc-sport-endpoint", "localhost:9001", "gRPC Sports server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Printf("failed running api server: %s\n", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*grpcRacingEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	//Register new endpoint for sports
	if err := sports.RegisterSportsHandlerFromEndpoint(
		ctx,
		mux,
		*grpcSportsEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	log.Printf("API server listening on: %s\n", *apiEndpoint)

	return http.ListenAndServe(*apiEndpoint, mux)
}
