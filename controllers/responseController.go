package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError_response(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": 0,
		"error":   "Database Error",
		"data":    nil,
		"len":     nil,
	})
}
