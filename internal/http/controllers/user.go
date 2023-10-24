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
	"k8s.io/apimachinery/pkg/util/validation"
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
			IsOwner:   elem.IsOwner,
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

func GetAuthorizedUser(c *gin.Context) {
	cr, err := time.Parse(time.RFC3339, c.GetString("user_created_at"))
	if err != nil {
		panic(err)
	}
	up, err := time.Parse(time.RFC3339, c.GetString("user_updated_at"))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, u.UserWithTokenResponse{Name: c.GetString("user_name"), IsOwner: c.GetBool("user_is_owner"), AccessToken: c.GetString("user_token"), CreatedAt: cr, UpdatedAt: up})
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
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	cr, err := time.Parse(time.RFC3339, elem.CreatedAt)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	up, err := time.Parse(time.RFC3339, elem.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusOK, u.UserResponse{Name: elem.Name, CreatedAt: cr, UpdatedAt: up})
}

func CreateUser(c *gin.Context) {
	var request u.CreateUserRequest
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
		IsOwner:  request.IsOwner,
	}
	err = user.CreateAccessToken(&model)
	if err != nil {
		panic(err)
	}
	t := time.Now()
	model.IsOwner = request.IsOwner
	model.CreatedAt = t.Format(time.RFC3339)
	model.UpdatedAt = t.Format(time.RFC3339)
	_, err = db.GetMongoClient().Database("registry").Collection("users").InsertOne(context.TODO(), model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusCreated, u.UserResponse{Name: model.Name, CreatedAt: t, UpdatedAt: t, IsOwner: request.IsOwner})
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
