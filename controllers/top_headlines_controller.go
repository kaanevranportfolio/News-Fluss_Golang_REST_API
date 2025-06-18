package controllers

import (
	"encoding/json"
	"net/http"

	"news-fluss/config"

	"github.com/gin-gonic/gin"
)

func (nc *NewsController) GetTopHeadlines(c *gin.Context) {
	cfg := config.GetConfig()
	apiKey := cfg.NewsAPIKey
	baseURL := cfg.NewsAPIBaseURL

	// You can add query parameters as needed, e.g., country=us
	req, err := http.NewRequest("GET", baseURL+"/top-headlines?country=us", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	// Add API key as header for NewsAPI
	req.Header.Add("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch news from NewsAPI"})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (nc *NewsController) GetTopHeadlinesByCategory(c *gin.Context) {
	cfg := config.GetConfig()
	apiKey := cfg.NewsAPIKey
	baseURL := cfg.NewsAPIBaseURL

	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category parameter is required"})
		return
	}

	req, err := http.NewRequest("GET", baseURL+"/top-headlines"+"?category="+category, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Add("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch news from NewsAPI"})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, result)
}