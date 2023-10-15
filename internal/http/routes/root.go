// Package routes provide the REST API route handlers
package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// InitRouter will return HTTP gin router
func InitRouter(ui bool, webAppPath string) *gin.Engine {
	router := gin.Default()
	if ui {
		router.Use(static.Serve("/", static.LocalFile(webAppPath, false)))
		router.NoRoute(func(c *gin.Context) {
			c.File(webAppPath + "/index.html")
		})
	}
	v1 := router.Group("/api/v1")
	initSetupRoutes(v1)
	initAuthRoutes(v1)
	initUsersRoutes(v1)

	return router
}
