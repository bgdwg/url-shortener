package service

import (
	"context"
	"errors"
	"log"
	"os"
	pb "url-shortener/protos"
	"url-shortener/storage"
)

var (
	InfoLogger  = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

type Server struct {
	Storage storage.Storage
}

func (s *Server) CreateUrl(ctx context.Context, req *pb.Url) (*pb.Key, error) {
	url := req.GetUrl()
	if url == "" {
		err := errors.New("no url in request")
		ErrorLogger.Println("CreateUrl: %s", err.Error())
		return &pb.Key{}, err
	}

	InfoLogger.Println("CreateUrl: %s", url)

	key, err := s.Storage.PutUrl(ctx, storage.Url(url))
	if err != nil {
		ErrorLogger.Println("Storage.PutUrl: %s", err.Error())
		return &pb.Key{}, err
	}

	return &pb.Key{ Key: string(key) }, nil
}

func (s *Server) GetUrl(ctx context.Context, req *pb.Key) (*pb.Url, error) {
	key := req.GetKey()
	if key == "" {
		err := errors.New("no key in request")
		ErrorLogger.Println("GetUrl: %s", err.Error())
		return &pb.Url{}, err
	}

	InfoLogger.Println("GetUrl: %s", key)

	url, err := s.Storage.GetUrl(ctx, storage.Key(key))
	if err != nil {
		ErrorLogger.Println("Storage.GetUrl: %s", err.Error())
		return &pb.Url{}, err
	}

	return &pb.Url{ Url: string(url) }, nil
}