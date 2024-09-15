package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NewNewNews/NewNews-Gateway/internal/auth"
	"github.com/NewNewNews/NewNews-Gateway/internal/database"
	"github.com/NewNewNews/NewNews-Gateway/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	db     *database.Database
	jwt    *auth.JWTManager
	logger zerolog.Logger
}

func New(db *database.Database, jwt *auth.JWTManager, logger zerolog.Logger) *Handler {
	return &Handler{db: db, jwt: jwt, logger: logger}
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

	w.WriteHeader(http.StatusCreated)
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
