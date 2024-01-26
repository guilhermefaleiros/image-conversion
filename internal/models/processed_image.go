package models

import (
	"github.com/google/uuid"
	"time"
)

type ProcessedImage struct {
	ID             string    `json:"id"`
	OriginalName   string    `json:"original_name"`
	StoredFileName string    `json:"stored_file_name"`
	RequestID      string    `json:"request_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewProcessedImage(OriginalName string, RequestID string) *ProcessedImage {
	id := uuid.New().String()
	return &ProcessedImage{
		ID:             id,
		OriginalName:   OriginalName,
		StoredFileName: id + "-" + OriginalName,
		RequestID:      RequestID,
		CreatedAt:      time.Now(),
	}
}
