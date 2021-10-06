package postgresql

import (
	"context"
	"fmt"
	"os"
	"testing"
	"url-shortener/internal/storage"
)

var (
	dbName      = os.Getenv("POSTGRES_DATABASE_NAME")
	userName    = os.Getenv("POSTGRES_USER_NAME")
	password    = os.Getenv("POSTGRES_PASSWORD")
	pgPort      = os.Getenv("POSTGRES_PORT")
	pgHost      = os.Getenv("POSTGRES_HOST")
	s = NewStorage(fmt.Sprintf("postgres://%s:%s@%s%s/%s",
		userName, password, pgHost, pgPort, dbName))
)

func TestStorage(t *testing.T) {
	t.Run("put non-existing URL", func(t *testing.T) {
		url := storage.URL("https://www.google.com/")
		_, err := s.PutURL(context.Background(), url)
		if err != nil {
			t.Errorf("PutURL failed: %v", err)
		}
	})

	t.Run("get by existing key", func(t *testing.T) {
		key := storage.Key("HMR5f91Qth")
		_, err := s.GetURL(context.Background(), key)
		if err != nil {
			t.Errorf("GetURL failed: %v", err)
		}
	})

	t.Run("put existing URL", func(t *testing.T) {
		url := storage.URL("https://www.google.com/")
		key, err := s.PutURL(context.Background(), url)
		if err != nil {
			t.Errorf("PutURL failed: %v", err)
		}
		if key != "HMR5f91Qth" {
			t.Errorf("PutURL failed: key must not change")
		}
	})

	t.Run("get by non-existing key", func(t *testing.T) {
		key := storage.Key("non_exists")
		_, err := s.GetURL(context.Background(), key)
		if err == nil {
			t.Error("GetURL failed: must return err != nil")
		}
	})
}
