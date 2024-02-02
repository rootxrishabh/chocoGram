package server

import (
	"log"
	"os"

	"github.com/rootxrishabh/chocoGram/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "fmt"
)

// Server function housing all routing and configuration
func Server() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.SocialNetwork(router)
	router.Run(":" + os.Getenv("PORT"))
}
