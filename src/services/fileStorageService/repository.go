package filestorageservice

type FileStorageRepository interface {
	GetFileURL(name string) (string, error)
	Upload(file string) (string, error)
	Delete(name string) (string, error)
}
