package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/gin-gonic/gin"
)

type AudioRequest struct {
    NewsID string `json:"id"`
}

// GetAudioFile retrieves an audio file for a given news ID
func (h *Handler) GetAudioFile(c *gin.Context) {
	// Get parameter from JSON request
	var req AudioRequest
	
	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error().Err(err).Msg("Failed to parse request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	url := req.NewsID

	newsID := c.Param("id")

	// Call the audio service to get the audio file
	fmt.Println("newsID:", newsID)
	fmt.Println("url:", url)

	audioResponse, err := h.newsVoice.GetAudioFile(context.Background(), &proto.AudioRequest{NewsId: url})
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get audio file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get audio file"})
		return
	}

	// Set the appropriate headers for file download
	c.Header("Content-Disposition", "attachment; filename="+audioResponse.FileName)
	c.Data(http.StatusOK, "audio/mpeg", audioResponse.AudioData)
}

// ReceiveNewsContent receives news content and processes it (presumably for text-to-speech conversion)
func (h *Handler) ReceiveNewsContent(c *gin.Context) {
	var req proto.NewsContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error().Err(err).Msg("Invalid news content data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news content data"})
		return
	}

	// Call the audio service to process the news content
	response, err := h.newsVoice.ReceiveNewsContent(context.Background(), &req)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to process news content")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process news content"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// StreamAudioFile streams an audio file for a given news ID (optional method, not in the original proto)
func (h *Handler) StreamAudioFile(c *gin.Context) {
	newsID := c.Param("id")

	// Call the audio service to get the audio file
	audioResponse, err := h.newsVoice.GetAudioFile(context.Background(), &proto.AudioRequest{NewsId: newsID})
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get audio file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get audio file"})
		return
	}

	// Set the appropriate headers for audio streaming
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Content-Disposition", "inline; filename="+audioResponse.FileName)

	// Create a reader from the audio data
	reader := bytes.NewReader(audioResponse.AudioData)

	// Stream the audio data
	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to stream audio file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stream audio file"})
		return
	}
}
