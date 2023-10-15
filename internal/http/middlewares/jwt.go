// Package middlewares provide every http middleware
package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/db"
	"github.com/resourcemod/registry/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

func HasToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		token := c.GetHeader("Authorization")
		if len(strings.TrimSpace(token)) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization header is required.")
			return
		}

		form := strings.Split(token, " ")
		if form[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Bearer token is not provided.")
			return
		}

		res := db.GetMongoClient().Database("registry").Collection("users").FindOne(context.TODO(), bson.D{
			{"access_token", form[1]},
		})
		if res.Err() != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res.Err().Error())
			return
		}
		var model models.User
		err := res.Decode(&model)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		r, err := time.Parse(time.RFC3339, model.ExpiredAt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		if r.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Token expired.")
			return
		}

		c.Next() // request
		// after request
	}
}
