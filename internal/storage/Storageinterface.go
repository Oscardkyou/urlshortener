// storage/storage.go
package storage

type StorageInterface interface {
	Save(shortURL, longURL string) error
	Resolve(shortURL string) (string, error)
}
