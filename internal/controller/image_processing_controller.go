package controller

import (
	"github.com/gin-gonic/gin"
	"image-conversor/internal/controller/effects"
	"image-conversor/internal/models"
	"image-conversor/internal/provider"
	"image-conversor/internal/repository"
	"net/http"
	"path/filepath"
	"strings"
)

type ProcessImageRequest struct {
	ImageID string   `json:"image_id"`
	Effects []string `json:"effects"`
}

type ImageProcessingResult struct {
	OriginalImageId  string   `json:"image_id"`
	Status           string   `json:"status"`
	Effects          []string `json:"effects"`
	ProcessedImageId string   `json:"processed_image_id"`
}

type ImageProcessingController struct {
	storageProvider                  *provider.LocalStorageProvider
	imageRepository                  *repository.ImageRepository
	imageProcessingRequestRepository *repository.ImageProcessingRequestRepository
	processedImageRepository         *repository.ProcessedImageRepository
	jpegEffect                       *effects.JPEGEffect
	sepiaEffect                      *effects.SepiaEffect
	grayscaleEffect                  *effects.GrayscaleEffect
	pngEffect                        *effects.PNGEffect
	invertColorsEffect               *effects.InvertColorsEffect
}

func replaceExtension(filePath, newExtension string) string {
	extension := filepath.Ext(filePath)
	if !strings.HasPrefix(newExtension, ".") {
		newExtension = "." + newExtension
	}
	return filePath[:len(filePath)-len(extension)] + newExtension
}

func (i *ImageProcessingController) applyEffects(effectType string, bytes []byte) ([]byte, string, error) {
	if effectType == "jpeg" {
		return i.jpegEffect.Apply(bytes)
	}
	if effectType == "grayscale" {
		return i.grayscaleEffect.Apply(bytes)
	}
	if effectType == "sepia" {
		return i.sepiaEffect.Apply(bytes)
	}
	if effectType == "png" {
		return i.pngEffect.Apply(bytes)
	}
	if effectType == "invert_colors" {
		return i.invertColorsEffect.Apply(bytes)
	}
	return bytes, "", nil
}

func buildErrorResponse(input ProcessImageRequest) ImageProcessingResult {
	return ImageProcessingResult{
		OriginalImageId:  input.ImageID,
		Status:           "error",
		ProcessedImageId: "",
		Effects:          input.Effects,
	}
}

func (i *ImageProcessingController) processRequest(input ProcessImageRequest) ImageProcessingResult {
	request := models.NewImageProcessingRequest(input.ImageID)
	for _, effect := range input.Effects {
		request.AddEffect(effect)
	}

	err := i.imageProcessingRequestRepository.Save(request)
	if err != nil {
		return buildErrorResponse(input)
	}

	img, err := i.imageRepository.FindById(request.ImageId)
	if err != nil {
		return buildErrorResponse(input)
	}

	newFile, err := i.storageProvider.Get(img.StoredFileName)
	format := filepath.Ext(img.OriginalName)

	if err != nil {
		return buildErrorResponse(input)
	}

	for _, effect := range request.Effects {
		newFile, format, err = i.applyEffects(effect.Effect, newFile)
		if err != nil {
			return buildErrorResponse(input)
		}
	}

	processedImage := models.NewProcessedImage(replaceExtension(img.OriginalName, format), request.ID)
	err = i.storageProvider.Save(processedImage.StoredFileName, newFile)
	if err != nil {
		return buildErrorResponse(input)
	}
	err = i.processedImageRepository.Save(processedImage)
	if err != nil {
		return buildErrorResponse(input)
	}
	return ImageProcessingResult{
		OriginalImageId:  input.ImageID,
		Status:           "success",
		Effects:          input.Effects,
		ProcessedImageId: processedImage.ID,
	}
}

func (i *ImageProcessingController) ProcessBatch(r *gin.Context) {
	var requests []ProcessImageRequest

	err := r.ShouldBindJSON(&requests)

	for _, request := range requests {
		for _, effect := range request.Effects {
			if !Contains(ValidEffects, effect) {
				r.JSON(http.StatusBadRequest, gin.H{"error": "invalid effect"})
				return
			}
		}
	}
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results := make(chan ImageProcessingResult)
	for _, request := range requests {
		go func(req ProcessImageRequest) {
			result := i.processRequest(req)
			results <- result
		}(request)
	}

	var processedImages []ImageProcessingResult

	for range requests {
		result := <-results
		processedImages = append(processedImages, result)
	}

	close(results)

	r.JSON(http.StatusOK, gin.H{"processed_images": processedImages})
}

func NewImageProcessingController(
	storageProvider *provider.LocalStorageProvider,
	imageRepository *repository.ImageRepository,
	imageProcessingRequestRepository *repository.ImageProcessingRequestRepository,
	processedImageRepository *repository.ProcessedImageRepository) *ImageProcessingController {
	return &ImageProcessingController{
		storageProvider:                  storageProvider,
		imageRepository:                  imageRepository,
		imageProcessingRequestRepository: imageProcessingRequestRepository,
		processedImageRepository:         processedImageRepository,
		sepiaEffect:                      effects.NewSepiaEffect(),
		grayscaleEffect:                  effects.NewGrayscaleEffect(),
		jpegEffect:                       effects.NewJPEGEffect(),
		pngEffect:                        effects.NewPNGEffect(),
		invertColorsEffect:               effects.NewInvertColorsEffect(),
	}
}
