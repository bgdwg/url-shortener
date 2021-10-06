package shortener

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"
	"url-shortener/internal/storage"
	"url-shortener/internal/storage/postgresql"

	pb "url-shortener/internal/pkg/proto"
)

var (
	dbName      = os.Getenv("POSTGRES_DATABASE")
	userName    = os.Getenv("POSTGRES_USER")
	password    = os.Getenv("POSTGRES_PASSWORD")
	pgAddr      = os.Getenv("POSTGRES_ADDR")
	s = postgresql.NewStorage(fmt.Sprintf("postgres://%s:%s@%s/%s",
		userName, password, pgAddr, dbName))
)

// copy: https://stackoverflow.com/questions/42102496/testing-a-grpc-service

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	pb.RegisterURLShortenerServiceServer(server, &Server{Storage: s})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestServer_CreateURL(t *testing.T) {
	tests := []struct {
		name    string
		url		string
		res     *pb.CreateURLResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"invalid request with empty URL",
			"",
			nil,
			codes.Unknown,
			fmt.Sprintf("parse \"%s\": empty url", ""),
		},
		{
			"invalid request with non-empty URL",
			"invalid_url.xyz",
			nil,
			codes.Unknown,
			fmt.Sprintf("parse \"%s\": invalid URI for request", "invalid_url.xyz"),
		},
		{
			"valid request with non-empty URL",
			"https://www.google.com/",
			&pb.CreateURLResponse{Key: "HMR5f91Qth"},
			codes.OK,
			"",
		},
		{
			"valid request with non-empty URL",
			"https://www.youtube.com/",
			&pb.CreateURLResponse{Key: "UOFDv2qx5y"},
			codes.OK,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewURLShortenerServiceClient(conn)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := &pb.CreateURLRequest{Url: test.url}
			response, err := client.CreateURL(ctx, request)
			if response != nil {
				if len(response.GetKey()) != len(test.res.GetKey()) {
					t.Error("response: expected", test.res.GetKey(), "received", response.GetKey())
				}
			}

			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != test.errCode {
						t.Error("error code: expected", test.errCode, "received", er.Code())
					}
					if er.Message() != test.errMsg {
						t.Error("error message: expected", test.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}

func TestServer_GetURL(t *testing.T) {
	tests := []struct {
		name    string
		key		string
		res     *pb.GetURLResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"invalid request with banned symbols",
			"+983821234",
			nil,
			codes.Unknown,
			fmt.Sprintf("invalid key requested"),
		},
		{
			"invalid request with incorrect key length",
			"123456789ABC",
			nil,
			codes.Unknown,
			fmt.Sprintf("invalid key requested"),
		},
		{
			"valid request with existing key",
			"HMR5f91Qth",
			&pb.GetURLResponse{Url: "https://www.google.com/"},
			codes.Unknown,
			"",
		},
		{
			"valid request with non-existing key",
			"non_exists",
			&pb.GetURLResponse{Url: "https://www.google.com/"},
			codes.Unknown,
			fmt.Sprintf("no URL with key %s - %v", "non_exists", storage.NotFoundError),
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewURLShortenerServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.GetURLRequest{Key: tt.key}
			response, err := client.GetURL(ctx, request)
			if response != nil {
				if response.GetUrl() != tt.res.GetUrl() {
					t.Error("response: expected", tt.res.GetUrl(), "received", response.GetUrl())
				}
			}

			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", tt.errCode, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}