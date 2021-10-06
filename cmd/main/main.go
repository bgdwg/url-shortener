package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	pb "url-shortener/internal/pkg/proto"
	"url-shortener/internal/server/shortener"
	"url-shortener/internal/storage/postgresql"
)

var (
	grpcPort    = os.Getenv("GRPC_PORT")
	dbName      = os.Getenv("POSTGRES_DATABASE")
	userName    = os.Getenv("POSTGRES_USER")
	password    = os.Getenv("POSTGRES_PASSWORD")
	pgAddr      = os.Getenv("POSTGRES_ADDR")
	dbURL = fmt.Sprintf("postgres://%s:%s@%s/%s", userName, password, pgAddr, dbName)
)

func main() {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterURLShortenerServiceServer(grpcServer, &shortener.Server{
			Storage: postgresql.NewStorage(dbURL),
	})
	log.Printf("start serving on %s", grpcPort)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
