package filestorageservice

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const (
	timeout       = 10 * time.Second
	containerName = "resumes"
)

type azureFileStorage struct {
	client *azblob.Client
}

func NewAzureFileStorage(accountName string, accountKey string) FileStorageRepository {
	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		panic(err)
	}

	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)

	client, err := azblob.NewClientWithSharedKeyCredential(accountURL, cred, nil)
	if err != nil {
		panic(err)
	}

	return &azureFileStorage{
		client: client,
	}
}

func (m *azureFileStorage) GetFileURL(name string) (string, error) {
	return "GetFileURL", nil
}

func (m *azureFileStorage) Upload(file io.Reader, filename string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := m.client.UploadStream(ctx, containerName, filename, file, nil)
	if err != nil {
		panic(err)
	}
	log.Println(res)

	return "Upload", nil
}

func (m *azureFileStorage) Delete(name string) (string, error) {
	return "Delete", nil
}
