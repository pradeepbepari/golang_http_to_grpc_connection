package handlers

import (
	"http_server/models"
	"http_server/service"
	"net/http"
	authentication "sdk-helper/authenticate"
	"sdk-helper/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	service service.AuthService
	logger  *logger.Logger
}

func NewAuthHandler(service service.AuthService, logger *logger.Logger) AuthHandler {
	return authHandler{service: service, logger: logger}
}
func (h authHandler) Login(c *gin.Context) {
	var request models.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password cannot be empty"})
		return
	}

	response, err := h.service.LoginUser(c, request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := authentication.GenerateJwtToken(response.ID, response.Name, response.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
func (a authHandler) Logout(c *gin.Context) {
	c.Header("Authorization", "")
	c.Request.Header.Set("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"message": "Signed out successfully"})
}
