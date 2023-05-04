package filestorageservice

type TestFileStorage struct{}

func NewTestFileStorage() FileStorageRepository {
	return &TestFileStorage{}
}

func (m *TestFileStorage) FindOne(name int) (string, error) {
	return "FindOne", nil
}

func (m *TestFileStorage) FindMany(names []string) (string, error) {
	return "FindMany", nil
}

func (m *TestFileStorage) Insert(file string) (string, error) {
	return "Insert", nil
}

func (m *TestFileStorage) Delete(name int) (string, error) {
	return "Delete", nil
}
