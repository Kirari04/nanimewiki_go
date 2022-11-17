package main

import (
	"ch/kirari/animeApi/controllers"
	"ch/kirari/animeApi/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/ping", controllers.Default)
		}
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	router.Run()
}
