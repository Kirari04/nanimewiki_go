package controllers

import (
	"ch/kirari/animeApi/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MapRequest_ListAnime struct {
	Index uint `uri:"index"`
}

func ListAnime(c *gin.Context) {
	// VALIDATOR
	var req MapRequest_ListAnime
	if c.ShouldBindUri(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
			"error":   "Bad Request",
			"data":    nil,
			"len":     nil,
		})
		return
	}

	// LOGIC
	limit, _ := strconv.Atoi(os.Getenv("api_lenlist"))
	offset := int(req.Index) * limit

	var animes []models.Anime
	models.DB.Limit(limit).Offset(offset).Find(&animes)
	var itemLen int64
	models.DB.Model(&models.Anime{}).Count(&itemLen)

	c.JSON(http.StatusOK, gin.H{
		"success": 1,
		"error":   "",
		"data":    animes,
		"len":     itemLen,
	})
}
