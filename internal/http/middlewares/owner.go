// Package middlewares provide every http middleware
package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/db"
	u "github.com/resourcemod/registry/pkg/api"
	"github.com/resourcemod/registry/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

func IsOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		if len(c.GetString("user_name")) != 0 {
			if !c.GetBool("user_is_owner") {
				c.AbortWithStatusJSON(http.StatusForbidden, u.UnauthorizedResponse{Message: "Forbidden.", Code: http.StatusForbidden})
				return
			}
			c.Next()
			return
		}

		token := c.GetHeader("Authorization")
		if len(strings.TrimSpace(token)) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: "Invalid token", Code: http.StatusUnauthorized})
			return
		}

		form := strings.Split(token, " ")
		if form[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: "Invalid token", Code: http.StatusUnauthorized})
			return
		}

		res := db.GetMongoClient().Database("registry").Collection("users").FindOne(context.TODO(), bson.D{
			{"access_token", form[1]},
		})
		if res.Err() != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: "Invalid token", Code: http.StatusUnauthorized})
			return
		}
		var model models.User
		err := res.Decode(&model)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: "Invalid token", Code: http.StatusUnauthorized})
			return
		}
		r, err := time.Parse(time.RFC3339, model.ExpiredAt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: err.Error(), Code: http.StatusUnauthorized})
			return
		}
		if r.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: "Invalid token", Code: http.StatusUnauthorized})
			return
		}

		if !model.IsOwner {
			c.AbortWithStatusJSON(http.StatusForbidden, u.UnauthorizedResponse{Message: "Forbidden.", Code: http.StatusForbidden})
			return
		}

		c.Set("user_is_owner", model.IsOwner)
		c.Next() // request
		// after request
	}
}
