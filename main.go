package main

import (
	"news-fluss/config"
	"news-fluss/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig() // Load configuration at startup

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run() // Start the server
}
