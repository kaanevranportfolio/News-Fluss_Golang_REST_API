package routes

import (
	"github.com/gin-gonic/gin"
	"news-fluss/controllers"
)

func SetupRoutes(router *gin.Engine) {
	newsController := controllers.NewsController{}

	router.GET("/everything", newsController.GetEverything)
	
}