// storage/interface.go
package storage

type StorageInterface interface {
	Save(URL string) string
	Resolve(id string) (string, error)
}
