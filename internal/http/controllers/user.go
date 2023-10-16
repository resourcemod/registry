// Package controllers provide http request handlers
package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/resourcemod/registry/internal/db"
	"github.com/resourcemod/registry/internal/services/user"
	u "github.com/resourcemod/registry/pkg/api"
	"github.com/resourcemod/registry/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

func GetUsersList(c *gin.Context) {
	collection := db.GetMongoClient().Database("registry").Collection("users")
	var users []u.UserResponse

	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.User
		err = cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		cr, err := time.Parse(time.RFC3339, elem.CreatedAt)
		if err != nil {
			panic(err)
		}
		up, err := time.Parse(time.RFC3339, elem.UpdatedAt)
		if err != nil {
			panic(err)
		}
		users = append(users, u.UserResponse{
			Name:      elem.Name,
			CreatedAt: cr,
			UpdatedAt: up,
		})
	}

	if err = cur.Err(); err != nil {
		panic(err)
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusOK, u.UsersListResponse{Users: users})
}

func GetUserByName(c *gin.Context) {
	var request u.GetUserByNameParams
	request.Name = strings.ToLower(c.Param("name"))
	res := db.GetMongoClient().Database("registry").Collection("users").FindOne(context.TODO(), bson.D{{
		"name", request.Name,
	}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	var elem models.User
	err := res.Decode(&elem)
	if err != nil {
		panic(err)
	}
	cr, err := time.Parse(time.RFC3339, elem.CreatedAt)
	if err != nil {
		panic(err)
	}
	up, err := time.Parse(time.RFC3339, elem.UpdatedAt)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, u.UserResponse{Name: elem.Name, CreatedAt: cr, UpdatedAt: up})
}

func CreateUser(c *gin.Context) {
	var request u.CreateUserRequest
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

func DeleteUser(c *gin.Context) {
	var request u.DeleteUserParams
	request.Name = strings.ToLower(c.Param("name"))
	res := db.GetMongoClient().Database("registry").Collection("users").FindOneAndDelete(context.TODO(), bson.D{{
		"name", request.Name,
	}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, u.DeleteUserResponse{Message: "Deleted"})
}
