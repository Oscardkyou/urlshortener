package shortener

import (
	"errors"
	"fmt"
	"hash/crc32"
	"urlshortener/storage"
)

type ShortenerService struct {
	store storage.StorageInterface
}

func NewShortenerService(store storage.StorageInterface) *ShortenerService {
	return &ShortenerService{store: store}
}

func (s *ShortenerService) Shorten(longURL string) (string, error) {
	shortURL := generateShortKey(longURL)
	err := s.store.Save(shortURL, longURL)
	if err != nil {
		return "", errors.New("failed to save the URL")
	}
	return shortURL, nil
}

func (s *ShortenerService) Expand(shortURL string) (string, error) {
	return s.store.Resolve(shortURL)
}

func generateShortKey(url string) string {
	return fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(url)))
}
