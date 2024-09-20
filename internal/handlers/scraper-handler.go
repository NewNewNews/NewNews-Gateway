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
func (h *Handler) UpdateNews(c *gin.Context) {
	var updateReq struct {
		ID        string `json:"id"`
		Data      string `json:"data"`
		Category  string `json:"category"`
		Date      string `json:"date"`
		Publisher string `json:"publisher"`
		URL       string `json:"url"`
	}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	resp, err := h.newsClient.UpdateNews(c, &proto.UpdateNewsRequest{
		Id:        updateReq.ID,
		Data:      updateReq.Data,
		Category:  updateReq.Category,
		Date:      updateReq.Date,
		Publisher: updateReq.Publisher,
		Url:       updateReq.URL,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"message": resp.Message,
	})
}

func (h *Handler) DeleteNews(c *gin.Context) {
	var deleteReq struct {
		ID string `json:"id"`
	}

	if err := c.ShouldBindJSON(&deleteReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	resp, err := h.newsClient.DeleteNews(c, &proto.DeleteNewsRequest{
		Id: deleteReq.ID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"message": resp.Message,
	})
}

func (h *Handler) GetOneNews(c *gin.Context) {
	var oneNewsReq struct {
		ID string `json:"id"`
	}

	if err := c.ShouldBindJSON(&oneNewsReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	resp, err := h.newsClient.GetOneNews(c, &proto.GetOneNewsRequest{
		Id: oneNewsReq.ID,
	})

	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get one news")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get one news"})
		return
	}

	c.JSON(http.StatusOK, resp.News)
}
