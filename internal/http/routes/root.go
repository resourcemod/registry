// Package routes provide the REST API route handlers
package routes

import (
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"time"
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
	if len(os.Getenv("IS_DEV")) > 0 {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "http://localhost:3000"
			},
			MaxAge: 12 * time.Hour,
		}))
	}
	v1 := router.Group("/api/v1")

	initSetupRoutes(v1)
	initAuthRoutes(v1)
	initUsersRoutes(v1)
	initContentRoutes(v1)
	initIntegrationsRoutes(v1)

	return router
}
