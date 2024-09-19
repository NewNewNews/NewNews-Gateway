package main

import (
	"log"

	"github.com/NewNewNews/NewNews-Gateway/internal/auth"
	"github.com/NewNewNews/NewNews-Gateway/internal/config"
	"github.com/NewNewNews/NewNews-Gateway/internal/database"
	"github.com/NewNewNews/NewNews-Gateway/internal/handlers"
	"github.com/NewNewNews/NewNews-Gateway/internal/logger"
	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.New()

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Disconnect()

	jwt := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpirationHours)

	var retryPolicy = `{
		"methodConfig": [{
			"name": [{"service": "grpc.examples.echo.Echo"}],
			"retryPolicy": {
				"MaxAttempts": 4,
				"InitialBackoff": ".01s",
				"MaxBackoff": ".01s",
				"BackoffMultiplier": 1.0,
				"RetryableStatusCodes": [ "UNAVAILABLE" ]
			}
		}]
	}`
	newsConn, err := grpc.NewClient(cfg.ScraperURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))

	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to news service")
	} else {
		// Log if connection is established
		logger.Info().Msg("Successfully connected to news service")
	}
	defer newsConn.Close()
	newsClient := proto.NewNewsServiceClient(newsConn)

	handler := handlers.New(db, jwt, logger, newsClient)

	// Initialize Gin
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Your frontend origin
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Define routes
	r.POST("/api/register", handler.Register)
	r.POST("/api/login", handler.Login)
	r.GET("/api/protected", auth.GinMiddleware(jwt), handler.Protected)
	r.GET("/api/news", handler.GetNews)
	r.POST("/api/scrape", handler.ScrapeNews)
	r.GET("/api/getall", handler.GetAllUsers)
	r.PUT("/api/user/update", handler.UpdateUserByEmail)
	r.DELETE("/api/user/remove", handler.DeleteUser)

	logger.Info().Msgf("Server starting on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
