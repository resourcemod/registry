// Package routes provide the REST API route handlers
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/http/controllers"
	"github.com/resourcemod/registry/internal/http/middlewares"
)

func initIntegrationsRoutes(rg *gin.RouterGroup) {
	rg.GET("/integrations", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetIntegrations)
	rg.POST("/integrations", middlewares.HasToken(), middlewares.IsOwner(), controllers.CreateIntegration)
	rg.DELETE("/integrations/:name", middlewares.HasToken(), middlewares.IsOwner(), controllers.DeleteIntegration)
	rg.GET("/integrations/:name/repositories", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetRepositories)
}
