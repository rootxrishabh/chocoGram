package server

import (
	"os"

	"github.com/rootxrishabh/chocoGram/routes"
	"github.com/gin-gonic/gin"
)

// Server function housing all routing and configuration
func Server() {
	router := gin.Default()
	routes.SocialNetwork(router)
	router.Run(":" + os.Getenv("PORT"))
}
