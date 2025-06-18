package controllers

import (
	"encoding/json"
	"net/http"
	"news-fluss/config"
	"news-fluss/models"

	"github.com/gin-gonic/gin"
)

type SourcesController struct{}

func (sc *SourcesController) GetSources(c *gin.Context) {
	cfg := config.GetConfig()
	apiKey := cfg.NewsAPIKey
	baseURL := cfg.NewsAPIBaseURL

	req, err := http.NewRequest("GET", baseURL+"/top-headlines/sources", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Add("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sources"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch sources from NewsAPI"})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (sc *SourcesController) GetSourcesByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	cfg := config.GetConfig()
	apiKey := cfg.NewsAPIKey
	baseURL := cfg.NewsAPIBaseURL

	req, err := http.NewRequest("GET", baseURL+"/top-headlines/sources?category="+category, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Add("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sources"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch sources from NewsAPI"})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (sc *SourcesController) SearchSource(c *gin.Context) {
	var source models.Source

	if err := c.ShouldBindJSON(&source); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get API key and base URL from config
	cfg := config.GetConfig()
	apiKey := cfg.NewsAPIKey
	baseURL := cfg.NewsAPIBaseURL

	// Build query parameters from source fields
	reqURL := baseURL + "/top-headlines/sources?"
	params := ""

	if source.Category != "" {
		params += "category=" + source.Category + "&"
	}
	if source.Language != "" {
		params += "language=" + source.Language + "&"
	}
	if source.Country != "" {
		params += "country=" + source.Country + "&"
	}
	if len(params) > 0 && params[len(params)-1] == '&' {
		params = params[:len(params)-1]
	}
	reqURL += params

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Add("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sources"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch sources from NewsAPI"})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, result)
}
