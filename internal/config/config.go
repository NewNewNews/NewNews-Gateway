package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port               string
	DatabaseURL        string
	JWTSecret          string
	JWTExpirationHours time.Duration
	ScraperURL         string
}

func Load() (*Config, error) {
	jwtExpirationHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:               os.Getenv("PORT"),
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		JWTExpirationHours: time.Duration(jwtExpirationHours) * time.Hour,
		ScraperURL:         os.Getenv("SCRAPER_SERVICE"),
	}, nil
}
