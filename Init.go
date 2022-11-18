package main

import (
	"ch/kirari/animeApi/controllers"
	"ch/kirari/animeApi/models"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	models.ConnectDatabase(os.Getenv("database"))
	var doSeed, _ = strconv.ParseBool(os.Getenv("database_seed"))
	if doSeed {
		models.SeedDatabase()
	}

	gin.SetMode(os.Getenv("gin_mode"))
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("trusted_proxie")})
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			anime := v1.Group("/anime")
			{
				anime.GET("/list", controllers.ListAnime)
				anime.GET("/list/:index", controllers.ListAnime)
			}
		}
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": 0,
			"error":   "Not Found",
			"data":    nil,
			"len":     nil,
		})
	})

	router.Run(os.Getenv("host") + ":" + os.Getenv("port"))
}
