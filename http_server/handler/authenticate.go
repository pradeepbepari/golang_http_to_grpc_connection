package handlers

import (
	"http_server/service"
	"sdk-helper/logger"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
	logger  *logger.Logger
}

func NewAuthHandler(service service.AuthService, logger *logger.Logger) AuthHandler {
	return authHandler{service: service, logger: logger}
}
func (h authHandler) Login(c *gin.Context) {
	// Implement login logic here
	c.JSON(200, gin.H{"message": "login successful"})
}
