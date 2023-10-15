// Package routes provide the REST API route handlers
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/http/controllers"
)

func initAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", controllers.Register)
	rg.POST("/login", controllers.Login)
}
