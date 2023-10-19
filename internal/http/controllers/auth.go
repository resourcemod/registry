// Package controllers provide http request handlers
package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/db"
	user "github.com/resourcemod/registry/internal/services/user"
	u "github.com/resourcemod/registry/pkg/api"
	"github.com/resourcemod/registry/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"k8s.io/apimachinery/pkg/util/validation"
	"net/http"
	"strings"
	"time"
)

func Register(c *gin.Context) {
	var request u.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	request.Name = strings.ToLower(request.GetName())
	if len(validation.IsDNS1123Label(request.Name)) > 0 {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: "The name must comply with RFC 1123 Label Names standard.", Code: http.StatusUnprocessableEntity})
		return
	}
	if len(request.GetPassword()) <= 5 {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: "Password should be at least 6 characters long.", Code: http.StatusUnprocessableEntity})
		return
	}

	if request.GetPassword() != request.GetPasswordConfirmation() {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: "Passwords missmatch.", Code: http.StatusUnprocessableEntity})
		return
	}
	hash, err := user.HashPassword(request.Password)
	if err != nil {
		panic(err)
	}
	model := models.User{
		Name:     request.Name,
		Password: hash,
		IsOwner:  false,
	}
	err = user.CreateAccessToken(&model)
	if err != nil {
		panic(err)
	}
	t := time.Now()
	model.CreatedAt = t.Format(time.RFC3339)
	model.UpdatedAt = t.Format(time.RFC3339)
	_, err = db.GetMongoClient().Database("registry").Collection("users").InsertOne(context.TODO(), model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusCreated, u.UserWithTokenResponse{Name: model.Name, AccessToken: model.AccessToken, CreatedAt: t, UpdatedAt: t, IsOwner: false})
}

func Login(c *gin.Context) {
	var request u.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	request.Name = strings.ToLower(request.GetName())
	res := db.GetMongoClient().Database("registry").Collection("users").FindOne(context.TODO(), bson.D{{"name", request.Name}})
	if res.Err() != nil {
		c.JSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: res.Err().Error(), Code: http.StatusUnauthorized})
		return
	}

	var model models.User
	err := res.Decode(&model)
	if err != nil {
		c.JSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: err.Error(), Code: http.StatusUnauthorized})
		return
	}
	if model.AccessToken == "" {
		c.JSON(http.StatusUnauthorized, u.UnauthorizedResponse{Message: err.Error(), Code: http.StatusUnauthorized})
		return
	}

	err = user.CreateAccessToken(&model)
	if err != nil {
		panic(err)
	}
	t := time.Now()
	model.UpdatedAt = t.Format(time.RFC3339)
	update := bson.D{{"$set", bson.D{{"access_token", model.AccessToken}, {"expired_at", model.ExpiredAt}, {"updated_at", model.UpdatedAt}}}}
	_, err = db.GetMongoClient().Database("registry").Collection("users").UpdateOne(context.TODO(), bson.D{{"name", request.GetName()}}, update)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	cr, err := time.Parse(time.RFC3339, model.CreatedAt)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, u.UserWithTokenResponse{Name: model.Name, AccessToken: model.AccessToken, UpdatedAt: t, CreatedAt: cr, IsOwner: model.IsOwner})
}
