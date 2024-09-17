package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/NewNewNews/NewNews-Gateway/internal/auth"
	"github.com/NewNewNews/NewNews-Gateway/internal/database"
	"github.com/NewNewNews/NewNews-Gateway/internal/models"
	"github.com/NewNewNews/NewNews-Gateway/internal/proto"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	db         *database.Database
	jwt        *auth.JWTManager
	logger     zerolog.Logger
	newsClient proto.NewsServiceClient
}

func New(db *database.Database, jwt *auth.JWTManager, logger zerolog.Logger, newsClient proto.NewsServiceClient) *Handler {
	return &Handler{db: db, jwt: jwt, logger: logger, newsClient: newsClient}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.HashedPassword = string(hashedPassword)
	if err := h.db.CreateUser(r.Context(), &user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.db.GetUserByEmail(r.Context(), credentials.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.jwt.Generate(user.ID, user.IsAdmin)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) Protected(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("user").(*jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := (*claims)["user_id"].(string)
	log := &models.Log{
		UserID:    userID,
		Action:    "Accessed protected endpoint",
		Timestamp: time.Now(),
	}
	if err := h.db.CreateLog(r.Context(), log); err != nil {
		h.logger.Error().Err(err).Msg("Failed to create log")
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Access granted to protected resource"})
}

func (h *Handler) GetNews(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	date := r.URL.Query().Get("date")

	resp, err := h.newsClient.GetNews(context.Background(), &proto.GetNewsRequest{
		Category: category,
		Date:     date,
	})
	if err != nil {
		http.Error(w, "Failed to get news", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp.News)
}

func (h *Handler) ScrapeNews(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.newsClient.ScrapeNews(context.Background(), &proto.ScrapeNewsRequest{
		Url: req.URL,
	})
	if err != nil {
		http.Error(w, "Failed to scrape news", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": resp.Success})
}

func (h *Handler) UpdateNews(w http.ResponseWriter, r *http.Request) {
	var updateReq struct {
		ID        string `json:"id"`
		Data      string `json:"data"`
		Category  string `json:"category"`
		Date      string `json:"date"`
		Publisher string `json:"publisher"`
		URL       string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	resp, err := h.newsClient.UpdateNews(context.Background(), &proto.UpdateNewsRequest{
		Id:        updateReq.ID,
		Data:      updateReq.Data,
		Category:  updateReq.Category,
		Date:      updateReq.Date,
		Publisher: updateReq.Publisher,
		Url:       updateReq.URL,
	})
	if err != nil {
		http.Error(w, "Failed to update news", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": resp.Success,
		"message": resp.Message,
	})
}

func (h *Handler) DeleteNews(w http.ResponseWriter, r *http.Request) {
	var deleteReq struct {
		ID        string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&deleteReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	resp, err := h.newsClient.DeleteNews(context.Background(), &proto.DeleteNewsRequest{
		Id:        deleteReq.ID,
	})

	if err != nil {
		http.Error(w, "Failed to delete news", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": resp.Success,
		"message": resp.Message,
	})
}

func (h *Handler) GetOneNews(w http.ResponseWriter, r *http.Request) {
	var oneNewsReq struct {
		ID        string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&oneNewsReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

    resp, err := h.newsClient.GetOneNews(context.Background(), &proto.GetOneNewsRequest{
        Id: oneNewsReq.ID,
    })

	if err != nil {
		log.Printf("Failed to get one news: %v", err)
		http.Error(w, "Failed to get one news", http.StatusInternalServerError)
		return
	}
	

    json.NewEncoder(w).Encode(resp.News)
}
