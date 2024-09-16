package main

import (
	"log"
	"net/http"

	"github.com/NewNewNews/NewNews-Gateway/internal/auth"
	"github.com/NewNewNews/NewNews-Gateway/internal/config"
	"github.com/NewNewNews/NewNews-Gateway/internal/database"
	"github.com/NewNewNews/NewNews-Gateway/internal/handlers"
	"github.com/NewNewNews/NewNews-Gateway/internal/logger"
	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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

	// Set up gRPC connection to news service
	newsConn, err := grpc.NewClient("news_service:50051", grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to news service")
	}
	defer newsConn.Close()
	newsClient := proto.NewNewsServiceClient(newsConn)

	handler := handlers.New(db, jwt, logger, newsClient)

	http.HandleFunc("/api/register", handler.Register)
	http.HandleFunc("/api/login", handler.Login)
	http.Handle("/api/protected", auth.Middleware(jwt, handler.Protected))
	http.HandleFunc("/api/news", handler.GetNews)
	http.HandleFunc("/api/scrape", handler.ScrapeNews)

	logger.Info().Msgf("Server starting on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
