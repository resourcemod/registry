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
	"net/http"
	"strings"
	"time"
)

func Register(c *gin.Context) {
	var request u.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	request.Name = strings.ToLower(request.GetName())

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
	}
	err = user.CreateAccessToken(&model)
	if err != nil {
		panic(err)
	}
	model.CreatedAt = time.Now().Format(time.RFC3339)
	model.UpdatedAt = time.Now().Format(time.RFC3339)
	_, err = db.GetMongoClient().Database("registry").Collection("users").InsertOne(context.TODO(), model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusCreated, u.CreateUserResponse{Name: model.Name, AccessToken: model.AccessToken})
}

func Login(c *gin.Context) {
	var request u.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	request.Name = strings.ToLower(request.GetName())
	res := db.GetMongoClient().Database("registry").Collection("users").FindOne(context.TODO(), bson.D{{"name", request.Name}})
	if res.Err() != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": res.Err().Error()})
		return
	}

	var model models.User
	err := res.Decode(&model)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if model.AccessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found."})
		return
	}

	err = user.CreateAccessToken(&model)
	if err != nil {
		panic(err)
	}
	model.UpdatedAt = time.Now().Format(time.RFC3339)
	update := bson.D{{"$set", bson.D{{"access_token", model.AccessToken}, {"expired_at", model.ExpiredAt}, {"updated_at", model.UpdatedAt}}}}
	_, err = db.GetMongoClient().Database("registry").Collection("users").UpdateOne(context.TODO(), bson.D{{"name", request.GetName()}}, update)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusCreated, u.CreateUserResponse{Name: model.Name, AccessToken: model.AccessToken})
}
