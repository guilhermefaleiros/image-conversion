package models

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	OriginalName   string    `json:"original_name"`
	StoredFileName string    `json:"stored_file_name"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewImage(OriginalName string) *Image {
	id := uuid.New().String()
	return &Image{
		ID:             id,
		OriginalName:   OriginalName,
		StoredFileName: id + "-" + OriginalName,
		CreatedAt:      time.Now(),
	}
}
