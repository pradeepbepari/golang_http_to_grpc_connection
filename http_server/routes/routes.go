package routes

import (
	handlers "http_server/handler"
	authentication "sdk-helper/authenticate"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	AuthHandler  handlers.AuthHandler
	UserHandler  handlers.UserHandler
	FileHandlers handlers.FileHandler
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func ApiRoutes(di Dependencies, router *gin.Engine) {
	routes := router.Group("/api")
	routes.POST("/register", di.UserHandler.RegisterUser)
	routes.POST("/login", di.AuthHandler.Login)

	routes.POST("/upload", di.FileHandlers.HandleFileUploader)
	routes.Use(authentication.Authenticate)
	{

	}
}
