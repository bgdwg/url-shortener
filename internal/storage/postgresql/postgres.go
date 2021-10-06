package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"url-shortener/internal/storage"
	"url-shortener/internal/utils/generator"
)

type PostgresStorage struct {
	Db *pgxpool.Pool
}

func NewStorage(dbURL string) *PostgresStorage {
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("unable to parse database URL - %v", err)
	}
	ctx := context.Background()
	db, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("unable to create connection pool - %v", err)
	}
	if _, err := db.Exec(ctx, `create table if not exists urls
	(id text primary key, url text not null)`); err != nil {
		log.Fatalf("unable to create table - %v", err)
	}
	return &PostgresStorage{Db: db}
}

func (s *PostgresStorage) PutURL(ctx context.Context, url storage.URL) (storage.Key, error) {
	var id string
	err := s.Db.QueryRow(ctx, "select id from urls where url=$1", url).Scan(&id)
	if err == pgx.ErrNoRows {
		for attempt := 1; attempt <= 5; attempt++ {
			id = generator.GetRandomKey()
			err = s.Db.QueryRow(ctx, "select id from urls where id=$1", id).Scan(&id)
			if err != pgx.ErrNoRows {
				continue
			}
			if _, err := s.Db.Exec(ctx, `insert into urls(id, url)
			values ($1, $2)`, id, url); err != nil {
				return "", fmt.Errorf("insertion error - %w", storage.BaseError)
			}
			return storage.Key(id), nil
		}
		return "", fmt.Errorf("insertion error - %w", storage.BaseError)
	}
	return storage.Key(id), nil
}

func (s *PostgresStorage) GetURL(ctx context.Context, key storage.Key) (storage.URL, error) {
	var url storage.URL
	err := s.Db.QueryRow(ctx, "select url from urls where id=$1", key).Scan(&url)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("no URL with key %v - %w", key, storage.NotFoundError)
		}
		return "", fmt.Errorf("selection error - %w", storage.BaseError)
	}
	return url, nil
}




