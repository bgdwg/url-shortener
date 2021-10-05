package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
	pb "url-shortener/protos"
	"url-shortener/server"
	"url-shortener/storage"
)

func main() {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = ":8081"
	}

	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUrlShortenerServer(grpcServer, &server.Server{
		Storage: storage.NewMemoryStorage(),
	})

	log.Printf("start serving on %s", grpcPort)

	if err = grpcServer.Serve(listener); err != nil {
		grpclog.Fatalf("failed to serve %v", err)
	}
}
