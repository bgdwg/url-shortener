package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
	pb "url-shortener/protos"
)

var (
	address = "localhost:8081"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewURLShortenerServiceClient(conn)
	// Contact the server and print out its response.
	var method, url, key string
	if len(os.Args) > 1 {
		method = os.Args[1]
		if method == "create" {
			url = os.Args[2]
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			createURLResp, err := c.CreateURL(ctx, &pb.CreateURLRequest{Url: url})
			if err != nil {
				log.Fatalf("could not create URL: %v", err)
			}
			log.Printf("Create short URL: %s", createURLResp.GetKey())
		} else if method == "get" {
			key = os.Args[2]
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			createURLResp, err := c.GetURL(ctx, &pb.GetURLRequest{Key: key})
			if err != nil {
				log.Fatalf("could not create URL: %v", err)
			}
			log.Printf("Create short URL: %s", createURLResp.GetUrl())
		} else {
			log.Fatalf("first argument must be 'create' or 'get': %v", err)
		}
	}
}



