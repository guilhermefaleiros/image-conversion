package controller

import (
	"github.com/gin-gonic/gin"
	"image-conversor/internal/repository"
	"net/http"
)

type ProcessedImageController struct {
	processedImageRepository *repository.ProcessedImageRepository
}

func (p *ProcessedImageController) Download(c *gin.Context) {
	id := c.Param("id")

	image, err := p.processedImageRepository.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}
	c.File("images/" + image.StoredFileName)
}

func NewProcessedImageController(processedImageRepository *repository.ProcessedImageRepository) *ProcessedImageController {
	return &ProcessedImageController{
		processedImageRepository,
	}
}
