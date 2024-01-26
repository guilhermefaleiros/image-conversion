package provider

import (
	"os"
	"path/filepath"
)

type LocalStorageProvider struct{}

func (p *LocalStorageProvider) Save(filePath string, file []byte) error {
	path := filepath.Join("images", filePath)
	err := os.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (p *LocalStorageProvider) Get(filePath string) ([]byte, error) {
	path := filepath.Join("images", filePath)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewLocalStorageProvider() *LocalStorageProvider {
	return &LocalStorageProvider{}
}
