package filestorageservice

import "io"

type FileStorageRepository interface {
	GetFileURL(name string) (string, error)
	Upload(file io.Reader, filename string) (string, error)
	Delete(name string) (string, error)
}
