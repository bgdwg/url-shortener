package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	pb "url-shortener/protos"
	"url-shortener/storage"
)

var (
	InfoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	validKey = regexp.MustCompile("^[a-zA-Z0-9_]*$")
)

type Server struct {
	pb.UnimplementedURLShortenerServiceServer
	Storage storage.Storage
}

func (s *Server) CreateURL(ctx context.Context, req *pb.CreateURLRequest) (*pb.CreateURLResponse, error) {
	resp := &pb.CreateURLResponse{}
	longURL := req.GetUrl()
	if _, err := url.ParseRequestURI(longURL); err != nil {
		ErrorLogger.Printf("CreateURL: invalid url requested: %v", err)
		return resp, err
	}
	InfoLogger.Printf("CreateURL: create short URL from: %s", longURL)
	key, err := s.Storage.PutURL(ctx, storage.URL(longURL))
	if err != nil {
		ErrorLogger.Printf("CreateURL: %v", err)
		return resp, err
	}
	resp.Key = string(key)
	return resp, nil
}

func (s *Server) GetURL(ctx context.Context, req *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	resp := &pb.GetURLResponse{}
	key := req.GetKey()
	fmt.Println(key)
	if !validKey.MatchString(key) {
		err := errors.New("invalid key requested")
		ErrorLogger.Printf("GetURL: %v", err)
		return resp, err
	}
	InfoLogger.Printf("GetURL: lookup for URL with key: %s", key)
	longURL, err := s.Storage.GetURL(ctx, storage.Key(key))
	if err != nil {
		ErrorLogger.Printf("GetURL: %v", err)
		return resp, err
	}
	resp.Url = string(longURL)
	return resp, nil
}

