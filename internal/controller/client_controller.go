package controller

import (
	"github.com/gin-gonic/gin"
	"image-conversor/internal/auth"
	"image-conversor/internal/models"
	"image-conversor/internal/repository"
	"net/http"
)

type CreateClientRequestDTO struct {
	Name string `json:"name"`
}

type GenerateAccessTokenRequestDTO struct {
	ClientId  string `json:"client_id"`
	AccessKey string `json:"access_key"`
}

type ClientController struct {
	clientRepository *repository.ClientRepository
}

func (c *ClientController) GenerateAccessToken(r *gin.Context) {
	var request GenerateAccessTokenRequestDTO
	err := r.ShouldBindJSON(&request)

	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if request.ClientId == "" || request.AccessKey == "" {
		r.JSON(http.StatusBadRequest, gin.H{"error": "client_id and access_key are required"})
		return
	}

	client, err := c.clientRepository.FindById(request.ClientId)

	if err != nil {
		r.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}

	if client.AccessKey != request.AccessKey {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid access key"})
		return
	}

	accessToken, err := auth.GenerateToken(client.ID)

	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": "error on generate access token"})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
	return
}

func (c *ClientController) Create(r *gin.Context) {
	var request CreateClientRequestDTO
	err := r.ShouldBindJSON(&request)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" {
		r.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	client := models.NewClient(request.Name)
	err = c.clientRepository.Save(client)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	r.JSON(http.StatusOK, gin.H{
		"id":         client.ID,
		"access_key": client.AccessKey,
		"name":       client.Name,
	})
	return
}

func NewClientController(clientRepository *repository.ClientRepository) *ClientController {
	return &ClientController{
		clientRepository,
	}
}
