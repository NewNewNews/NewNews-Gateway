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

	audioConn, err := grpc.NewClient(cfg.VoiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))

	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to audio service")
	} else {
		logger.Info().Msg("Successfully connected to audio service")
	}
	defer audioConn.Close()
	audioClient := proto.NewAudioServiceClient(audioConn)

	summaryConn, err := grpc.NewClient(cfg.SummaryURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))

	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to summary service")
	} else {
		logger.Info().Msg("Successfully connected to summary service")
	}
	defer summaryConn.Close()
	summaryClient := proto.NewSummaryServiceClient(summaryConn)

	handler := handlers.New(db, jwt, logger, newsClient, audioClient, summaryClient)

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
	// Route user
	r.POST("/api/register", handler.Register)
	r.POST("/api/login", handler.Login)
	r.POST("/api/logout", handler.Logout)
	// r.GET("/api/me", auth.AuthMiddleware(jwt), handler.GetMe)
	protected := r.Group("/api/protected")
	protected.Use(auth.AuthMiddleware(jwt))
	{
		protected.GET("/me", handler.GetMe)
	}

	admin := r.Group("/api/admin")
	admin.Use(auth.AuthMiddleware(jwt), auth.AdminMiddleware())
	{
		admin.GET("/getall", handler.GetAllUsers)
	}

	r.GET("/api/protected", auth.GinMiddleware(jwt), handler.Protected)
	// r.GET("/api/getall", handler.GetAllUsers)
	r.PUT("/api/user/update", handler.UpdateUserByEmail)
	r.DELETE("/api/user/remove", handler.DeleteUser)

	// Route scraper
	r.GET("/api/news", handler.GetNews)
	r.POST("/api/scrape", handler.ScrapeNews)
	r.PUT("/api/news", handler.UpdateNews)
	r.DELETE("/api/news", handler.DeleteNews)
	r.GET("/api/news/one", handler.GetOneNews)

	// Route audio
	audioGroup := r.Group("api/audio")
	{
		audioGroup.GET("/:id", handler.GetAudioFile)
		audioGroup.POST("/content", handler.ReceiveNewsContent)
		audioGroup.GET("/stream/:id", handler.StreamAudioFile)
	}

	// Route summary
	summaryGroup := r.Group("api/summary")
	{
		summaryGroup.GET("/one", handler.SummarizeNews)
	}

	// admin := r.Group("/api/admin")
	// admin.Use(auth.AuthMiddleware(jwt), auth.AdminMiddleware())
	// {
	// 	admin.POST("/scrape", handler.ScrapeNews)
	// 	admin.PUT("/news", handler.UpdateNews)
	// 	admin.DELETE("/news", handler.DeleteNews)
	// 	admin.GET("/users", handler.GetAllUsers)
	// 	admin.PUT("/user", handler.UpdateUserByEmail)
	// 	admin.DELETE("/user", handler.DeleteUser)
	// }

	logger.Info().Msgf("Server starting on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
