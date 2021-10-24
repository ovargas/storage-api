package driver

import "context"

type Properties map[string]interface{}

type Provider interface {
	Store(ctx context.Context, filename string, content []byte) (string, error)
	Download(ctx context.Context, filename string) ([]byte, error)
	Delete(ctx context.Context, filename string) error
}

type Driver interface {
	Initialize(cfg Properties) (Provider, error)
}
