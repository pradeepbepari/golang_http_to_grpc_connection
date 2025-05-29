package handlers

import (
	"http_server/models"
	"http_server/service"
	"sdk-helper/logger"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
	logger  *logger.Logger
}

func NewHandler(service service.UserService, logger *logger.Logger) UserHandler {
	return userHandler{
		service: service,
		logger:  logger,
	}
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object} map[string]string
// @Failure      500  {object} map[string]string
// @Router       /accounts/{id} [get]
func (h userHandler) RegisterUser(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.service.CreateUser(c, request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user created"})
}
