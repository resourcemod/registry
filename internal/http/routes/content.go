// Package routes provide the REST API route handlers
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/http/controllers"
	"github.com/resourcemod/registry/internal/http/middlewares"
)

func initContentRoutes(rg *gin.RouterGroup) {
	rg.GET("/content/:type", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetContentList)
	rg.GET("/content/:type/:name", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetContentByName)
	//rg.GET("/content/:type/:name/download", middlewares.HasToken(), controllers.GetDownloadLink)
	rg.POST("/content", middlewares.HasToken(), middlewares.IsOwner(), controllers.CreateContent)
	rg.PUT("/content/:type/:name", middlewares.HasToken(), middlewares.IsOwner(), controllers.UpdateContent)
	rg.DELETE("/content/:type/:name", middlewares.HasToken(), middlewares.IsOwner(), controllers.DeleteContent)
}
