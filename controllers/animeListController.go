package controllers

import (
	"ch/kirari/animeApi/models"
	"ch/kirari/animeApi/setups"
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
		CustomError_response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// LOGIC
	limit, _ := strconv.Atoi(os.Getenv("api_lenlist"))
	offset := int(req.Index) * limit

	var animes []models.Anime
	setups.DB.Limit(limit).Offset(offset).Find(&animes)
	var itemLen int64
	setups.DB.Model(&models.Anime{}).Count(&itemLen)

	OK_response(c, animes, itemLen)
}
