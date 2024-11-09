package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/gin-gonic/gin"
)

// ComparisonRequest represents the request body for getting comparisons.
type ComparisonRequest struct {
	NewsID string `json:"news_id"`
}

func (h *Handler) GetComparison(c *gin.Context) {
	// Parse the JSON request to get the news_id
	var req ComparisonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error().Err(err).Msg("Failed to parse request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	comparisonResponse, err := h.newsCompare.GetComparison(
		context.Background(),
		&proto.GetComparisonRequest{NewsId: req.NewsID},
	)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get comparison")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comparison"})
		return
	}

	// Log the comparison entries for debugging
	for _, entry := range comparisonResponse.Entries {
		fmt.Printf("Key: %s, Values: %v\n", entry.Key, entry.Values)
	}

	// Return the comparison data as JSON response
	c.JSON(http.StatusOK, comparisonResponse)
}
