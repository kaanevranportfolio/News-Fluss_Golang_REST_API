package routes

import (
	"github.com/gin-gonic/gin"
	"news-fluss/controllers"

)

func SetupRoutes(router *gin.Engine) {
	newsController := controllers.NewsController{}
	sourcesController := controllers.SourcesController{}

	router.GET("/news", newsController.GetNews)
	router.POST("/news", newsController.SearchNews)

	router.GET("/topHeadlines", newsController.GetTopHeadlines)
	router.GET("/topHeadlines/:category", newsController.GetTopHeadlinesByCategory)

	router.GET("/sources", sourcesController.GetSources)
	router.GET("/sources/:category", sourcesController.GetSourcesByCategory)
	router.POST("/sources", sourcesController.SearchSource)
}