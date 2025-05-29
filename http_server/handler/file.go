package handlers

import (
	"http_server/service"
	"sdk-helper/logger"

	"github.com/gin-gonic/gin"
)

type fileHandler struct {
	service service.FileService
	logger  *logger.Logger
}

func NewFileHandler(service service.FileService, logger *logger.Logger) fileHandler {
	return fileHandler{
		service: service,
		logger:  logger,
	}
}
func (h fileHandler) HandleFileUploader(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{"error": "method invalid header request"})
		return
	}
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(400, gin.H{"error": "max file size should be 10mb"})
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "failed to get file from header"})
		return
	}
	defer file.Close()
	filekey := header.Filename
	url, err := h.service.FileUploadToS3(c, file, filekey)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to upload file"})
		return
	}
	c.JSON(200, gin.H{"message": url})
}
