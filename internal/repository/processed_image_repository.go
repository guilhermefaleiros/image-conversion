package repository

import (
	"gorm.io/gorm"
	"image-conversor/internal/models"
)

type ProcessedImageRepository struct {
	db *gorm.DB
}

func (r *ProcessedImageRepository) Save(image *models.ProcessedImage) error {
	result := r.db.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProcessedImageRepository) FindById(id string) (*models.ProcessedImage, error) {
	var image models.ProcessedImage
	result := r.db.First(&image, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &image, nil
}

func NewProcessedImageRepository(db *gorm.DB) *ProcessedImageRepository {
	return &ProcessedImageRepository{
		db,
	}
}
