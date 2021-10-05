package storage

import (
	"context"
	"fmt"
	"sync"
	"url-shortener/generator"
)

type MemoryStorage struct {
	Urls map[Key]URL
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Urls: make(map[Key]URL),
	}
}

func (s *MemoryStorage) PutURL(_ context.Context, url URL) (Key, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for attempt := 0; attempt < 5; attempt++ {
		urlKey := Key(generator.GetRandomKey())
		if _, found := s.Urls[urlKey]; found {
			continue
		}
		s.Urls[urlKey] = url
		fmt.Println(s.Urls)
		return urlKey, nil
	}
	return "", fmt.Errorf("too much attempts during inserting - %w", CollisionError)
}

func (s *MemoryStorage) GetURL(_ context.Context, key Key) (URL, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, found := s.Urls[key]
	fmt.Println(s.Urls)
	if !found {
		return "", fmt.Errorf("no url with key %v - %w", key, NotFoundError)
	}
	return url, nil
}