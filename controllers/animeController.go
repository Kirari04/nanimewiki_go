package controllers

import (
	"ch/kirari/animeApi/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

// Default controller handles returning the hello world JSON response
func Default(c *gin.Context) {
	var animes []models.Anime
	models.DB.Find(&animes)
	c.JSON(http.StatusOK, gin.H{"message": animes})
}
