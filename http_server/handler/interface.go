package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	RegisterUser(c *gin.Context)
}
type FileHandler interface {
	HandleFileUploader(c *gin.Context)
}
