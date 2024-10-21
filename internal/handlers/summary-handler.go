package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/gin-gonic/gin"
)

type SummaryRequest struct {
    Url string `json:"url"`
}

func (h *Handler) SummarizeNews(c *gin.Context) {
	
	// Get parameter from JSON request
	var req SummaryRequest
	
	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error().Err(err).Msg("Failed to parse request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call the summary service to get the summary file
	fmt.Println("url:", req.Url)

	summaryResponse, err := h.newsSummary.SummarizeNews(context.Background(), &proto.SummaryNewsRequest{Url: req.Url})
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get summary")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get summary"})
		return
	}

	fmt.Printf("Summary: "+summaryResponse.SummarizedText)
	c.JSON(http.StatusOK, summaryResponse)
}
