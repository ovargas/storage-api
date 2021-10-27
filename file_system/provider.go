package file_system

import (
	"context"
	"fmt"
	"github.com/ovargas/storage-api/storage"
	"github.com/ovargas/storage-api/storage/driver"
	"io/ioutil"
	"os"
	"path/filepath"
)

type local struct{}

func init() {
	storage.Register("FileSystem", local{})
}

func (local) Initialize(properties driver.Properties) (driver.Provider, error) {
	folder := properties["folder"].(string)

	_ = os.MkdirAll(folder, os.ModePerm)

	return &provider{
		folder: properties["folder"].(string),
	}, nil
}

type provider struct {
	folder string
}

func (p *provider) Store(ctx context.Context, filename string, content []byte) (string, error) {

	fullPath := fmt.Sprintf("%s/%s", p.folder, filename)

	path:= filepath.Dir(fullPath)
	_ = os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(fullPath)

	defer func() {
		_ = file.Close()
	}()

	if err != nil {
		return "", err
	}

	_, err = file.Write(content)

	if err != nil {
		return "", err
	}
	return filename, err
}

func (p *provider) Download(ctx context.Context, filename string) ([]byte, error) {
	file, err := os.Open(filename)
	defer func() {
		_ = file.Close()
	}()

	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}

func (p *provider) Delete(ctx context.Context, filename string) error {
	return os.Remove(filename)
}
