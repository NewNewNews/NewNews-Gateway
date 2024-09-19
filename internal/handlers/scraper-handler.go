package handlers

import (
	"net/http"

	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNews(c *gin.Context) {
	category := c.Query("category")
	date := c.Query("date")

	resp, err := h.newsClient.GetNews(c, &proto.GetNewsRequest{
		Category: category,
		Date:     date,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get news"})
		return
	}

	c.JSON(http.StatusOK, resp.News)
}

func (h *Handler) ScrapeNews(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.newsClient.ScrapeNews(c, &proto.ScrapeNewsRequest{
		Url: req.URL,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scrape news"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
