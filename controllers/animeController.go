package controllers

import (
	"ch/kirari/animeApi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAnime(c *gin.Context) {
	var animes []models.Anime
	models.DB.Find(&animes)

	var maxLen = 100
	var itemLen = len(animes)
	var maxIndex = itemLen / maxLen
	var strIndex, hasIndex = c.Params.Get("index")
	if !hasIndex {
		strIndex = "0"
	}
	intIndex, err := strconv.Atoi(strIndex)

	if err != nil || intIndex > maxIndex || intIndex < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
			"error":   "Bad Request",
			"data":    nil,
			"len":     nil,
		})
		return
	}

	var startIndex = intIndex * maxLen
	var endIndex = (intIndex + 1) * maxLen
	if itemLen < endIndex {
		endIndex = itemLen
	}

	c.JSON(http.StatusOK, gin.H{
		"success": 1,
		"error":   "",
		"data":    animes[startIndex:endIndex],
		"len":     itemLen,
	})
}
