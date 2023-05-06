package filestorageservice

import "io"

type testFileStorage struct{}

func NewTestFileStorage() FileStorageRepository {
	return &testFileStorage{}
}

func (m *testFileStorage) GetFileURL(name string) string {
	return "GetFileURL"
}

func (m *testFileStorage) Upload(file io.Reader) (string, error) {
	return "Upload", nil
}

func (m *testFileStorage) Delete(name string) (string, error) {
	return "Delete", nil
}
