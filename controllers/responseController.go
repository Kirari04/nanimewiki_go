package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError_response(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": 0,
		"error":   "Internal Server Error",
		"data":    nil,
		"len":     nil,
	})
}
func CustomError_response(c *gin.Context, errorCode int, errorMessage string) {
	c.JSON(errorCode, gin.H{
		"success": 0,
		"error":   errorMessage,
		"data":    nil,
		"len":     nil,
	})
}

func OK_response(c *gin.Context, data interface{}, len interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": 1,
		"error":   nil,
		"data":    data,
		"len":     len,
	})
}
