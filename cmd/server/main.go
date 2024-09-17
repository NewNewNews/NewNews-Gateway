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
	"github.com/rs/cors"
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
	newsConn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithInsecure()) //"news_service:50051"
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to news service")
	}
	defer newsConn.Close()
	newsClient := proto.NewNewsServiceClient(newsConn)

	handler := handlers.New(db, jwt, logger, newsClient)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Your frontend origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/api/register", handler.Register)
	mux.HandleFunc("/api/login", handler.Login)
	mux.Handle("/api/protected", auth.Middleware(jwt, handler.Protected))
	mux.HandleFunc("/api/news", handler.GetNews)
	mux.HandleFunc("/api/oneNews", handler.GetOneNews)
	mux.HandleFunc("/api/scrape", handler.ScrapeNews)
	mux.HandleFunc("/api/updateNews", handler.UpdateNews)
	mux.HandleFunc("/api/deleteNews", handler.DeleteNews)
	mux.HandleFunc("/api/getall", handler.GetAllUsers)
	mux.HandleFunc("/api/user/update", handler.UpdateUserByEmail)
	mux.HandleFunc("/api/user/remove", handler.DeleteUser)
	// Apply CORS middleware
	handlerWithCORS := corsHandler.Handler(mux)

	logger.Info().Msgf("Server starting on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handlerWithCORS); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
