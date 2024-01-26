package api

import (
	"github.com/gin-gonic/gin"
	"image-conversor/internal/auth"
	"image-conversor/internal/config"
	"image-conversor/internal/controller"
	"image-conversor/internal/provider"
	"image-conversor/internal/repository"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		isValid, _ := auth.ValidateToken(token)
		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}

func StartApiServer() {
	r := gin.Default()
	db := config.NewDatabase()

	clientRepository := repository.NewClientRepository(db)
	imageRepository := repository.NewImageRepository(db)
	imageProcessingRequestRepository := repository.NewImageProcessingRequestRepository(db)
	processedImageRepository := repository.NewProcessedImageRepository(db)

	clientController := controller.NewClientController(clientRepository)
	storageProvider := provider.NewLocalStorageProvider()
	imageController := controller.NewImageController(imageRepository, storageProvider)
	processedImageController := controller.NewProcessedImageController(processedImageRepository)
	imageProcessingController := controller.NewImageProcessingController(
		storageProvider,
		imageRepository,
		imageProcessingRequestRepository,
		processedImageRepository,
	)

	r.POST("/client", clientController.Create)
	r.POST("/generate-token", clientController.GenerateAccessToken)
	r.Use(AuthMiddleware())
	r.POST("/image", imageController.Upload)
	r.POST("/image/process", imageProcessingController.ProcessBatch)
	r.GET("/processed-image/:id/download", processedImageController.Download)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
