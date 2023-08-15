package filestorageservice

import "io"

type FileStorageRepository interface {
	GetFileURL(name string) string
	Upload(file io.Reader) (string, error)
	Delete(name string) (string, error)
}
