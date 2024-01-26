package repository

import (
	"gorm.io/gorm"
	"image-conversor/internal/models"
)

type ClientRepository struct {
	db *gorm.DB
}

func (r *ClientRepository) Save(image *models.Client) error {
	result := r.db.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ClientRepository) FindById(id string) (*models.Client, error) {
	var image models.Client
	result := r.db.First(&image, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &image, nil
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		db,
	}
}
