package handlers

import (
	"net/http"

	"github.com/NewNewNews/NewNews-Gateway/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.HashedPassword = string(hashedPassword)
	if err := h.db.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.db.GetUserByEmail(c, credentials.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := h.jwt.Generate(user.ID, user.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the token in a cookie
	c.SetCookie(
		"auth_token",
		token,
		int(h.jwt.GetExpiration().Seconds()), // Max age in seconds
		"/",                                  // Path
		"",                                   // Domain
		false,                                 // Secure
		false,                                 // HttpOnly
	)

	// c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"email": user.Email,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (h *Handler) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.db.GetUserByID(c, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.db.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUserByEmail(c *gin.Context) {
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if updatedUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	if err := h.db.UpdateUserByEmail(c, updatedUser.Email, &updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var requestBody struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if requestBody.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	if err := h.db.DeleteUser(c, requestBody.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(http.StatusOK)
}
