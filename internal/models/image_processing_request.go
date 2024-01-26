package models

import (
	"github.com/google/uuid"
	"time"
)

type ImageProcessingRequest struct {
	ID          string                         `json:"id"`
	ImageId     string                         `json:"image_id"`
	RequestDate time.Time                      `json:"request_date"`
	Effects     []ImageProcessingRequestEffect `gorm:"foreignKey:request_id" json:"operations"`
}

func (r *ImageProcessingRequest) AddEffect(effect string) {
	r.Effects = append(r.Effects, ImageProcessingRequestEffect{
		ID:        uuid.New().String(),
		RequestId: r.ID,
		Effect:    effect,
	})
}

type ImageProcessingRequestEffect struct {
	ID        string `json:"id"`
	RequestId string `json:"request_id"`
	Effect    string `json:"effect"`
}

func NewImageProcessingRequest(ImageId string) *ImageProcessingRequest {
	return &ImageProcessingRequest{
		ID:          uuid.New().String(),
		ImageId:     ImageId,
		RequestDate: time.Now(),
		Effects:     make([]ImageProcessingRequestEffect, 0),
	}
}
