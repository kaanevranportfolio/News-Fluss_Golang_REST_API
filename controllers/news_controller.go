package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"news-fluss/config"
	"news-fluss/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NewsController is a struct for news-related handlers.
type NewsController struct{}

func (nc *NewsController) GetNews(c *gin.Context) {
	cfg := config.GetConfig()
	apiKey := cfg.NewsAPIKey
	baseURL := cfg.NewsAPIBaseURL

	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	req, err := http.NewRequest("GET", baseURL+"/everything"+"?q="+query, nil)
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
		var apiError map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&apiError)
		c.JSON(resp.StatusCode, gin.H{
			"error":            "Failed to fetch news from NewsAPI",
			"newsapi_response": apiError,
		})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (nc *NewsController) SearchNews(c *gin.Context) {
	cfg := config.GetConfig()
	apiKey := cfg.NewsAPIKey
	baseURL := cfg.NewsAPIBaseURL

	var reqBody models.News

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Build query parameters using url.Values
	values := url.Values{}
	if reqBody.Q != "" {
		values.Set("q", reqBody.Q)
	}
	if reqBody.Sources != "" {
		values.Set("sources", reqBody.Sources)
	}
	if reqBody.Domains != "" {
		values.Set("domains", reqBody.Domains)
	}
	if reqBody.ExcludeDomains != "" {
		values.Set("excludeDomains", reqBody.ExcludeDomains)
	}
	if reqBody.From != "" {
		values.Set("from", reqBody.From)
	}
	if reqBody.To != "" {
		values.Set("to", reqBody.To)
	}
	if reqBody.Language != "" {
		values.Set("language", reqBody.Language)
	}
	if reqBody.SortBy != "" {
		values.Set("sortBy", reqBody.SortBy)
	}
	if reqBody.PageSize != 0 {
		values.Set("pageSize", strconv.Itoa(reqBody.PageSize))
	}
	if reqBody.Page != 0 {
		values.Set("page", strconv.Itoa(reqBody.Page))
	}
	if reqBody.SearchIn != "" {
		values.Set("searchIn", reqBody.SearchIn)
	}

	req, err := http.NewRequest("GET", baseURL+"/everything?"+values.Encode(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("NewsAPI request URL:", req.URL.String()) // Log the request URL
	req.Header.Add("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiError map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&apiError)
		c.JSON(resp.StatusCode, gin.H{
			"error":            "Failed to fetch news from NewsAPI",
			"newsapi_response": apiError,
		})
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
