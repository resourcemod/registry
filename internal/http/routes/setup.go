// Package routes provide the REST API route handlers
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/http/controllers"
)

func initSetupRoutes(rg *gin.RouterGroup) {
	rg.GET("/setup", controllers.GetSetupRequired)
	rg.POST("/setup/create", controllers.Setup)
}
