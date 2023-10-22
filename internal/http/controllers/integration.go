// Package controllers provide http request handlers
package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v56/github"
	"github.com/resourcemod/registry/internal/db"
	u "github.com/resourcemod/registry/pkg/api"
	"github.com/resourcemod/registry/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"k8s.io/apimachinery/pkg/util/validation"
	"net/http"
	"strings"
)

func CreateIntegration(c *gin.Context) {
	var request u.CreateIntegrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	if !c.GetBool("user_is_owner") {
		c.JSON(http.StatusForbidden, u.ForbiddenResponse{Message: "Forbidden", Code: http.StatusForbidden})
		return
	}

	request.Name = strings.ToLower(request.GetName())
	if len(validation.IsDNS1123Label(request.Name)) > 0 {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: "The name must comply with RFC 1123 Label Names standard.", Code: http.StatusUnprocessableEntity})
		return
	}

	client := github.NewClient(nil).WithAuthToken(request.GetAccessToken())
	_, _, err := client.Repositories.ListAll(context.TODO(), &github.RepositoryListAllOptions{})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	model := models.Integration{
		Name:        request.Name,
		Type:        models.TYPE_GIT,
		Host:        models.HOST_GITHUB,
		AccessToken: request.GetAccessToken(),
	}

	_, err = db.GetMongoClient().Database("registry").Collection("integration").InsertOne(context.TODO(), model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusCreated, u.IntegrationResponse{Name: model.Name, AccessToken: model.AccessToken, Host: model.Host, Type: u.IntegrationResponseTypeGit})
}

func GetIntegrations(c *gin.Context) {
	collection := db.GetMongoClient().Database("registry").Collection("integration")
	var integrations []u.IntegrationResponse

	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Integration
		err = cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		integrations = append(integrations, u.IntegrationResponse{
			Name:        elem.Name,
			Host:        elem.Host,
			Type:        u.IntegrationResponseTypeGit,
			AccessToken: elem.AccessToken,
		})
	}

	if err = cur.Err(); err != nil {
		panic(err)
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusOK, u.IntegrationsResponse{Integrations: integrations})
}

func DeleteIntegration(c *gin.Context) {
	var request u.DeleteIntegrationParams
	request.Name = strings.ToLower(c.Param("name"))
	res := db.GetMongoClient().Database("registry").Collection("integration").FindOneAndDelete(context.TODO(), bson.D{{
		"name", request.Name,
	}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, u.DeleteIntegrationResponse{Message: "Deleted"})
}
