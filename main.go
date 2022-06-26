package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ginMode, exists := os.LookupEnv("GIN_MODE")
	if !exists {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"*"}
	config.ExposeHeaders = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	router.Run(":" + port)

}
