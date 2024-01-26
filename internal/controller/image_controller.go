package controller

import (
	"github.com/gin-gonic/gin"
	"image-conversor/internal/models"
	"image-conversor/internal/provider"
	"image-conversor/internal/repository"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type ImageController struct {
	imageRepository      *repository.ImageRepository
	localStorageProvider *provider.LocalStorageProvider
}

var validImageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

func isValidImageExtension(fileName string) bool {
	extension := strings.ToLower(filepath.Ext(fileName))
	_, ok := validImageExtensions[extension]
	return ok
}

func (i *ImageController) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	if !isValidImageExtension(header.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension"})
		return
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	image := models.NewImage(header.Filename)
	err = i.localStorageProvider.Save(image.StoredFileName, bytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = i.imageRepository.Save(image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"id": image.ID,
	})
	return
}

func NewImageController(imageRepository *repository.ImageRepository, localStorageProvider *provider.LocalStorageProvider) *ImageController {
	return &ImageController{
		imageRepository,
		localStorageProvider,
	}
}
