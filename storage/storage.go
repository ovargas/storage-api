package storage

import (
	"context"
	"fmt"
	"github.com/ovargas/storage-api/storage/driver"
)

type Storage struct {
	provider driver.Provider
}

func New(driverName string, properties driver.Properties) (*Storage, error) {
	driversMu.RLock()
	driver, ok := drivers[driverName]
	driversMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("storage: unknown provider %q (forgotten import?)", driverName)
	}

	provider, err := driver.Initialize(properties)
	if err != nil {
		return nil, err
	}

	return &Storage{provider: provider}, nil
}

func (s *Storage) Store(ctx context.Context, filename string, content []byte) (string, error) {
	//TODO: Add logging or something
	return s.provider.Store(ctx, filename, content)
}

func (s *Storage) Download(ctx context.Context, filename string) ([]byte, error) {
	//TODO: Add logging or something
	return s.provider.Download(ctx, filename)
}

func(s *Storage) Delete(ctx context.Context, filename string) error {
	//TODO: Add logging or something
	return s.provider.Delete(ctx, filename)
}
