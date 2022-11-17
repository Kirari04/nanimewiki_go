package controllers

import (
	"ch/kirari/animeApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {
	var animes []models.Anime
	models.DB.Find(&animes)
	c.JSON(http.StatusOK, gin.H{"message": animes})
}
