package filestorageservice

type TestFileStorage struct{}

func NewTestFileStorage() FileStorageRepository {
	return &TestFileStorage{}
}

func (m *TestFileStorage) GetFileURL(name string) (string, error) {
	return "GetFileURL", nil
}

func (m *TestFileStorage) Insert(file string) (string, error) {
	return "Insert", nil
}

func (m *TestFileStorage) Delete(name string) (string, error) {
	return "Delete", nil
}
