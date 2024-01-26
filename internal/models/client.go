package models

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/google/uuid"
)

type Client struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	AccessKey string `json:"access_key"`
}

func GenerateRandomKey(keySize int) (string, error) {
	bytes := make([]byte, keySize)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func NewClient(Name string) *Client {
	accessKey, err := GenerateRandomKey(32)
	if err != nil {
		panic(err)
	}
	return &Client{
		ID:        uuid.New().String(),
		Name:      Name,
		AccessKey: accessKey,
	}
}
