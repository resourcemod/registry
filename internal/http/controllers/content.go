// Package controllers provide http request handlers
package controllers

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

func GetContentList(c *gin.Context) {
	contentType := c.Param("type")
	var t u.ContentResponseType
	if contentType == models.TYPE_PLUGIN {
		t = u.ContentResponseTypePlugin
	} else {
		t = u.ContentResponseTypeExtension
	}
	collection := db.GetMongoClient().Database("registry").Collection("content")
	var content []u.ContentResponse

	cur, err := collection.Find(context.TODO(), bson.D{{"type", contentType}})
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Content
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
		content = append(content, u.ContentResponse{
			Name:        elem.Name,
			Version:     elem.Version,
			Type:        t,
			Description: elem.Description,
			IsPublic:    elem.IsPublic,
			UserName:    elem.UserName,
			CreatedAt:   cr,
			UpdatedAt:   up,
		})
	}

	if err = cur.Err(); err != nil {
		panic(err)
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusOK, u.ContentListResponse{Content: content})
}

func GetContentByName(c *gin.Context) {
	contentType := c.Param("type")
	var t u.ContentResponseType
	if contentType == models.TYPE_PLUGIN {
		t = u.ContentResponseTypePlugin
	} else {
		t = u.ContentResponseTypeExtension
	}
	name := strings.ToLower(c.Param("name"))
	collection := db.GetMongoClient().Database("registry").Collection("content")

	res := collection.FindOne(context.TODO(), bson.D{{"type", contentType}, {"name", name}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}
	var elem models.Content
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

	c.JSON(http.StatusOK, u.ContentResponse{
		Name:        elem.Name,
		Version:     elem.Version,
		Type:        t,
		Description: elem.Description,
		IsPublic:    elem.IsPublic,
		UserName:    elem.UserName,
		CreatedAt:   cr,
		UpdatedAt:   up,
	})
}

func CreateContent(c *gin.Context) {
	var request u.UploadContentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	request.Name = strings.ToLower(request.GetName())
	var contentType string
	var contentResponseType u.ContentResponseType
	if request.GetType() == u.UploadContentRequestTypePlugin {
		contentType = models.TYPE_PLUGIN
		contentResponseType = u.ContentResponseTypePlugin
	} else {
		contentType = models.TYPE_EXTENSION
		contentResponseType = u.ContentResponseTypeExtension
	}
	//todo: add validation
	cr := time.Now()
	up := time.Now()
	model := models.Content{
		Name:        request.Name,
		Version:     request.GetVersion(),
		Type:        contentType,
		Description: request.GetDescription(),
		IsPublic:    request.GetIsPublic(),
		UserName:    c.GetString("user_name"), // from jwt middleware
		Repository:  request.GetRepository(),
		CreatedAt:   cr.Format(time.RFC3339),
		UpdatedAt:   up.Format(time.RFC3339),
	}
	_, err := db.GetMongoClient().Database("registry").Collection("content").InsertOne(context.TODO(), model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusCreated, u.ContentResponse{
		Name:        request.Name,
		Version:     request.Version,
		Type:        contentResponseType,
		Description: request.Description,
		IsPublic:    request.IsPublic,
		UserName:    model.UserName,
		CreatedAt:   cr,
		UpdatedAt:   up,
	})
}

func UpdateContent(c *gin.Context) {
	var request u.UpdateContentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	name := strings.ToLower(c.Param("name"))
	t := c.Param("type")
	var contentType u.ContentResponseType
	if t == models.TYPE_PLUGIN {
		contentType = u.ContentResponseTypePlugin
	} else {
		contentType = u.ContentResponseTypeExtension
	}
	collection := db.GetMongoClient().Database("registry").Collection("content")

	res := collection.FindOne(context.TODO(), bson.D{{"type", t}, {"name", name}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}
	var model models.Content
	err := res.Decode(&model)
	if err != nil {
		panic(err)
	}
	model.UpdatedAt = time.Now().Format(time.RFC3339)
	model.Description = request.Description
	model.Repository = request.Repository
	model.IsPublic = request.IsPublic
	model.Version = request.Version

	update := bson.D{{"$set", bson.D{
		{"updated_at", model.UpdatedAt},
		{"description", model.Description},
		{"repository", model.Repository},
		{"is_public", model.IsPublic},
		{"version", model.Version},
	}}}
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"name", name}, {"type", t}}, update)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	cr, err := time.Parse(time.RFC3339, model.CreatedAt)
	if err != nil {
		panic(err)
	}
	up, err := time.Parse(time.RFC3339, model.UpdatedAt)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, u.ContentResponse{
		Name:        model.Name,
		Version:     model.Version,
		Type:        contentType,
		Description: model.Description,
		IsPublic:    model.IsPublic,
		UserName:    model.UserName,
		CreatedAt:   cr,
		UpdatedAt:   up,
	})
}

func DeleteContent(c *gin.Context) {
	var request u.DeleteContentParams
	request.Name = strings.ToLower(c.Param("name"))
	if c.Param("type") == models.TYPE_PLUGIN {
		request.Type = u.DeleteContentTypePlugin
	} else {
		request.Type = u.DeleteContentTypeExtension
	}
	res := db.GetMongoClient().Database("registry").Collection("content").FindOneAndDelete(context.TODO(), bson.D{
		{"name", request.Name},
		{"type", c.Param("type")},
	})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, u.DeleteContentResponse{Message: "Deleted"})
}
