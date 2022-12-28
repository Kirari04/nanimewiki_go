package main

import (
	"ch/kirari/animeApi/console"
	"ch/kirari/animeApi/controllers"
	"ch/kirari/animeApi/setups"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"net/http"
)

func main() {
	// setup the enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	// setup the database connection
	setups.ConnectDatabase(os.Getenv("database"))

	// run console mode if required
	doConsole := flag.String("console", "false", "console")
	doSeed := flag.String("seed", "", "seed")
	flag.Parse()
	if *doConsole != "false" {
		console.Console(*doConsole, *doSeed)
		return
	}

	// configure gin
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

	// run gin
	router.Run(os.Getenv("host") + ":" + os.Getenv("port"))
}
