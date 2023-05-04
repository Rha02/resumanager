package filestorageservice

type TestFileStorage struct{}

func NewTestFileStorage() FileStorageRepository {
	return &TestFileStorage{}
}

func (m *TestFileStorage) GetFileURL(name string) (string, error) {
	return "GetFileURL", nil
}

func (m *TestFileStorage) Upload(file string) (string, error) {
	return "Upload", nil
}

func (m *TestFileStorage) Delete(name string) (string, error) {
	return "Delete", nil
}
