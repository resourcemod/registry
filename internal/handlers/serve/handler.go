// Package serve provide handle function to handle serve command
package serve

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/http/routes"
	"net/http"
)

func HandleServeCommand(ctx context.Context, host string, port string, static string, ui bool) error {
	engine := routes.InitRouter(ui, static)
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := engine.Run(host + ":" + port); err != nil {
		return err
	}
	return nil
}
