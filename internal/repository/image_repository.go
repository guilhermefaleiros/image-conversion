package repository

import (
	"gorm.io/gorm"
	"image-conversor/internal/models"
)

type ImageRepository struct {
	db *gorm.DB
}

func (r *ImageRepository) Save(image *models.Image) error {
	result := r.db.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ImageRepository) FindById(id string) (*models.Image, error) {
	var image models.Image
	result := r.db.First(&image, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &image, nil
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{
		db,
	}
}
