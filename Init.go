package main

import (
	"ch/kirari/animeApi/controllers"
	"ch/kirari/animeApi/setups"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	setups.ConnectDatabase(os.Getenv("database"))
	doSeed, err := strconv.ParseBool(os.Getenv("database_seed"))
	if err != nil {
		log.Panic("Failed to parse database_seed")
	}
	if doSeed {
		setups.SeedDatabase()
	}

	gin.SetMode(os.Getenv("gin_mode"))
	router := gin.New()
	if os.Getenv("gin_mode") == "debug" {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())
	router.SetTrustedProxies([]string{os.Getenv("trusted_proxie")})

	// ROUTES
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			anime := v1.Group("/anime")
			{
				anime.GET("/list", controllers.ListAnime)
				anime.GET("/list/:index", controllers.ListAnime)
			}

			user := v1.Group("/user")
			{
				user.POST("/register", controllers.UserRegister)
			}
		}
	}

	// ERROR ROUTES
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": 0,
			"error":   "Not Found",
			"data":    nil,
			"len":     nil,
		})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": 0,
			"error":   "Method not allowed",
			"data":    nil,
			"len":     nil,
		})
	})

	router.Run(os.Getenv("host") + ":" + os.Getenv("port"))
}
