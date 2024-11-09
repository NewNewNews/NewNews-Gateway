package handlers

import (
	"time"

	"github.com/NewNewNews/NewNews-Gateway/internal/auth"
	"github.com/NewNewNews/NewNews-Gateway/internal/database"
	"github.com/NewNewNews/NewNews-Gateway/internal/models"
	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

type Handler struct {
	db          *database.Database
	jwt         *auth.JWTManager
	logger      zerolog.Logger
	newsClient  proto.NewsServiceClient
	newsVoice   proto.AudioServiceClient
	newsSummary proto.SummaryServiceClient
	newsCompare proto.ComparisonServiceClient
}

func New(db *database.Database, jwt *auth.JWTManager, logger zerolog.Logger, newsClient proto.NewsServiceClient, newsVoice proto.AudioServiceClient, newsSummary proto.SummaryServiceClient, newsCompare proto.ComparisonServiceClient) *Handler {
	return &Handler{db: db, jwt: jwt, logger: logger, newsClient: newsClient, newsVoice: newsVoice, newsSummary: newsSummary, newsCompare: newsCompare}
}

func (h *Handler) Protected(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	standardClaims, ok := claims.(*jwt.StandardClaims)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid claims type"})
		return
	}

	userID := standardClaims.Subject

	log := &models.Log{
		UserID:    userID,
		Action:    "Accessed protected endpoint",
		Timestamp: time.Now(),
	}
	if err := h.db.CreateLog(c, log); err != nil {
		h.logger.Error().Err(err).Msg("Failed to create log")
	}

	c.JSON(200, gin.H{"message": "Access granted to protected resource"})
}
