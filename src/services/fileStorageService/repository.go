package filestorageservice

type FileStorageRepository interface {
	FindOne(name int) (string, error)
	FindMany(names []string) (string, error)
	Insert(file string) (string, error)
	Delete(name int) (string, error)
}
