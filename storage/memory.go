package storage

import (
	"context"
	"fmt"
	"sync"
	"url-shortener/generator"
)

type MemoryStorage struct {
	Urls map[Key]Url

	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Urls: make(map[Key]Url),
	}
}

func (s *MemoryStorage) PutUrl(_ context.Context, url Url) (Key, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for attempt := 0; attempt < 5; attempt++ {
		urlKey := Key(generator.GetRandomKey())
		if _, found := s.Urls[urlKey]; found {
			continue
		}
		s.Urls[urlKey] = url
		return urlKey, nil
	}

	return "", fmt.Errorf("too much attempts during inserting - %w", CollisionError)
}

func (s *MemoryStorage) GetUrl(_ context.Context, key Key) (Url, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, found := s.Urls[key]
	if !found {
		return "", fmt.Errorf("no url with key %v - %w", key, NotFoundError)
	}

	return url, nil
}