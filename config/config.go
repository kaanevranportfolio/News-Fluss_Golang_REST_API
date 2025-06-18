package config

import (
	"log"
	"os"
	"sync"
	"github.com/joho/godotenv"
)

type Config struct {
	NewsAPIKey     string
	NewsAPIBaseURL string
}

var (
	cfg  Config
	once sync.Once
)

func LoadConfig() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		cfg = Config{
			NewsAPIKey:     getEnv("NEWS_API_KEY", ""),
			NewsAPIBaseURL: getEnv("NEWSAPI_BASE_URL", ""),
		}
	})
}

func GetConfig() Config {
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
