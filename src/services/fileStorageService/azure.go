package filestorageservice

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/google/uuid"
)

const (
	timeout       = 10 * time.Second
	containerName = "resumes"
)

var accountURL string

type azureFileStorage struct {
	client *azblob.Client
}

func NewAzureFileStorage(accountName string, accountKey string) FileStorageRepository {
	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		panic(err)
	}

	accountURL = fmt.Sprintf("https://%s.blob.core.windows.net", accountName)

	client, err := azblob.NewClientWithSharedKeyCredential(accountURL, cred, nil)
	if err != nil {
		panic(err)
	}

	return &azureFileStorage{
		client: client,
	}
}

func (m *azureFileStorage) GetFileURL(name string) string {
	fileURL := fmt.Sprintf("%s/%s/%s", accountURL, containerName, name)
	return fileURL
}

func (m *azureFileStorage) Upload(file io.Reader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filename := uuid.New().String() + ".pdf"

	_, err := m.client.UploadStream(ctx, containerName, filename, file, nil)
	if err != nil {
		panic(err)
	}

	return filename, nil
}

func (m *azureFileStorage) Delete(name string) (string, error) {
	return "Delete", nil
}
