package filestorageservice

import (
	"errors"
	"io"
)

type testFileStorage struct{}

func NewTestFileStorage() FileStorageRepository {
	return &testFileStorage{}
}

func (m *testFileStorage) GetFileURL(name string) string {
	return "GetFileURL"
}

func (m *testFileStorage) Upload(file io.Reader) (string, error) {
	data, _ := io.ReadAll(file)
	if string(data) == "error" {
		return "", errors.New("error uploading file")
	}
	return "Upload", nil
}

func (m *testFileStorage) Delete(name string) (string, error) {
	return "Delete", nil
}
