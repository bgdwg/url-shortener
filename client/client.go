package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
	pb "url-shortener/protos"
)

const (
	address     = "localhost:8081"
	defaultName = ""
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUrlShortenerClient(conn)

	// Contact the server and print out its response.
	url := defaultName
	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp1, err := c.CreateUrl(ctx, &pb.Url{Url: url})
	if err != nil {
		log.Fatalf("could not create url: %v", err)
	}
	log.Printf("Key: %s", resp1.GetKey())

	resp2, err := c.GetUrl(ctx, &pb.Key{Key: resp1.GetKey()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp2.GetUrl())
}



