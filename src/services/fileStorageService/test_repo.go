package filestorageservice

import "io"

type testFileStorage struct{}

func NewTestFileStorage() FileStorageRepository {
	return &testFileStorage{}
}

func (m *testFileStorage) GetFileURL(name string) (string, error) {
	return "GetFileURL", nil
}

func (m *testFileStorage) Upload(file io.Reader, filename string) (string, error) {
	return "Upload", nil
}

func (m *testFileStorage) Delete(name string) (string, error) {
	return "Delete", nil
}
