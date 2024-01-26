package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"image-conversor/internal/models"
)

func NewDatabase() *gorm.DB {
	dsn := "host=localhost user=image-conversor password=image-conversor dbname=image-conversor-database port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.Image{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Client{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.ImageProcessingRequest{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.ImageProcessingRequestEffect{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.ProcessedImage{})
	if err != nil {
		return nil
	}
	return db
}
