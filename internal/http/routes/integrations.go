// Package routes provide the REST API route handlers
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/http/controllers"
	"github.com/resourcemod/registry/internal/http/middlewares"
)

func initIntegrationsRoutes(rg *gin.RouterGroup) {
	rg.GET("/integration", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetIntegrations)
	rg.POST("/integration", middlewares.HasToken(), middlewares.IsOwner(), controllers.CreateIntegration)
	rg.DELETE("/integration/:name", middlewares.HasToken(), middlewares.IsOwner(), controllers.DeleteIntegration)
}
