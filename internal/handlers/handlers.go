package handlers

import (
	"context"
	"encoding/json"
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

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// claims, ok := r.Context().Value("user").(*jwt.MapClaims)
	// if !ok || !(*claims)["is_admin"].(bool) {
	// 	http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
	// 	return
	// }

	users, err := h.db.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *Handler) UpdateUserByEmail(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure email is provided in the request body
	if updatedUser.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if err := h.db.UpdateUserByEmail(r.Context(), updatedUser.Email, &updatedUser); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Define a struct to capture the email from the request body
	var requestBody struct {
		Email string `json:"email"`
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Extract the email
	email := requestBody.Email
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Perform the delete operation
	if err := h.db.DeleteUser(r.Context(), email); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
