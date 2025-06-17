package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// NewsController is a struct for news-related handlers.
type NewsController struct{}

func (nc *NewsController) GetEverything(c *gin.Context) {

	// Load .env file from parent directory
	err := godotenv.Load(".env")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load .env file"})
		return
	}

	apiKey := os.Getenv("NEWS_API_KEY")
	baseURL := os.Getenv("NEWSAPI_BASE_URL")

	if apiKey == "" || baseURL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing NEWSAPI_KEY or NEWSAPI_BASE_URL in environment"})
		return
	}

	// You can add query parameters as needed, e.g., q=bitcoin
	req, err := http.NewRequest("GET", baseURL+"/everything"+"?q=bitcoin", nil)
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

