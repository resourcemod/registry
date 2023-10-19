// Package routes provide the REST API route handlers
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/http/controllers"
	"github.com/resourcemod/registry/internal/http/middlewares"
)

func initUsersRoutes(rg *gin.RouterGroup) {
	rg.GET("/users", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetUsersList)
	rg.GET("/user", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetAuthorizedUser)
	rg.GET("/users/:name", middlewares.HasToken(), middlewares.IsOwner(), controllers.GetUserByName)
	rg.POST("/users", middlewares.HasToken(), middlewares.IsOwner(), controllers.CreateUser)
	rg.DELETE("/users/:name", middlewares.HasToken(), middlewares.IsOwner(), controllers.DeleteUser)
}
