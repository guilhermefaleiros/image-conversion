package repository

import (
	"gorm.io/gorm"
	"image-conversor/internal/models"
)

type ImageProcessingRequestRepository struct {
	db *gorm.DB
}

func (r *ImageProcessingRequestRepository) Save(request *models.ImageProcessingRequest) error {
	result := r.db.Create(request)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ImageProcessingRequestRepository) FindById(id string) (*models.ImageProcessingRequest, error) {
	var request models.ImageProcessingRequest
	result := r.db.Preload("Effects").First(&request, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &request, nil
}

func NewImageProcessingRequestRepository(db *gorm.DB) *ImageProcessingRequestRepository {
	return &ImageProcessingRequestRepository{
		db,
	}
}
