package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"url-shortener/internal/pkg/proto"
	"url-shortener/internal/storage"
)

var (
	InfoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	validKey = regexp.MustCompile("^[a-zA-Z0-9_]*$")
)

type Server struct {
	proto.UnimplementedURLShortenerServiceServer
	Storage storage.Storage
}

func (s *Server) CreateURL(ctx context.Context, req *proto.CreateURLRequest) (*proto.CreateURLResponse, error) {
	longURL := req.GetUrl()
	if _, err := url.ParseRequestURI(longURL); err != nil {
		ErrorLogger.Printf("CreateURL: invalid url requested: %v", err)
		return nil, err
	}
	InfoLogger.Printf("CreateURL: create short URL from: %s", longURL)
	key, err := s.Storage.PutURL(ctx, storage.URL(longURL))
	if err != nil {
		ErrorLogger.Printf("CreateURL: %v", err)
		return nil, err
	}
	return &proto.CreateURLResponse{Key: string(key)}, nil
}

func (s *Server) GetURL(ctx context.Context, req *proto.GetURLRequest) (*proto.GetURLResponse, error) {
	key := req.GetKey()
	fmt.Println(key)
	if !validKey.MatchString(key) {
		err := errors.New("invalid key requested")
		ErrorLogger.Printf("GetURL: %v", err)
		return nil, err
	}
	InfoLogger.Printf("GetURL: lookup for URL with key: %s", key)
	longURL, err := s.Storage.GetURL(ctx, storage.Key(key))
	if err != nil {
		ErrorLogger.Printf("GetURL: %v", err)
		return nil, err
	}
	return &proto.GetURLResponse{Url: string(longURL)}, nil
}

