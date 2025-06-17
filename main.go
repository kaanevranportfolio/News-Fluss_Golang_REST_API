package main

import (
	"github.com/gin-gonic/gin"
	//"news-fluss/controllers"
	"news-fluss/routes"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(); err != nil {
		panic(err)
	}
}