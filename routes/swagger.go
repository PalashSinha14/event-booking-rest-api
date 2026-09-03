package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/palashsinha14/go-rest-api/docs"
)

// RegisterSwaggerRoute serves the interactive API docs at /swagger/index.html,
// reading the spec generated into docs/ by `swag init` (see main.go's
// @title/@version/... comment block for the general API info, and each
// handler's own godoc comment for its individual documentation).
func RegisterSwaggerRoute(server *gin.Engine) {
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
