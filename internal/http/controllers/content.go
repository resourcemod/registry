// Package controllers provide http request handlers
package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v56/github"
	"github.com/resourcemod/registry/internal/db"
	u "github.com/resourcemod/registry/pkg/api"
	"github.com/resourcemod/registry/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"k8s.io/apimachinery/pkg/util/validation"
	"net/http"
	"os"
	"path/filepath"
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
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	for cur.Next(context.TODO()) {
		var elem models.Content
		err = cur.Decode(&elem)
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

		content = append(content, u.ContentResponse{
			Name:        elem.Name,
			Type:        t,
			Description: elem.Description,
			IsPublic:    elem.IsPublic,
			UserName:    elem.UserName,
			Repository: u.Repository{
				FullName:    elem.Repository.FullName,
				GitURL:      elem.Repository.GitUrl,
				Integration: elem.Repository.Integration,
			},
			CreatedAt: cr,
			UpdatedAt: up,
		})
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

	c.JSON(http.StatusOK, u.ContentResponse{
		Name:        elem.Name,
		Type:        t,
		Description: elem.Description,
		IsPublic:    elem.IsPublic,
		UserName:    elem.UserName,
		Repository: u.Repository{
			FullName:    elem.Repository.FullName,
			GitURL:      elem.Repository.GitUrl,
			Integration: elem.Repository.Integration,
		},
		CreatedAt: cr,
		UpdatedAt: up,
	})
}

func CreateContent(c *gin.Context) {
	var request u.UploadContentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	request.Name = strings.ToLower(request.GetName())
	if len(validation.IsDNS1123Label(request.Name)) > 0 {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: "The name must comply with RFC 1123 Label Names standard.", Code: http.StatusUnprocessableEntity})
		return
	}
	var contentType string
	var contentResponseType u.ContentResponseType
	if request.GetType() == u.UploadContentRequestTypePlugin {
		contentType = models.TYPE_PLUGIN
		contentResponseType = u.ContentResponseTypePlugin
	} else {
		contentType = models.TYPE_EXTENSION
		contentResponseType = u.ContentResponseTypeExtension
	}

	cr := time.Now()
	up := time.Now()
	r := request.GetRepository()
	model := models.Content{
		Name:        request.Name,
		Type:        contentType,
		Description: request.GetDescription(),
		IsPublic:    request.GetIsPublic(),
		UserName:    c.GetString("user_name"), // from jwt middleware
		Repository: models.Repository{
			Integration: r.GetIntegration(),
			GitUrl:      r.GetGitURL(),
			FullName:    r.GetFullName(),
		},
		CreatedAt: cr.Format(time.RFC3339),
		UpdatedAt: up.Format(time.RFC3339),
	}

	res := db.GetMongoClient().Database("registry").Collection("integrations").FindOne(context.TODO(), bson.D{{
		"name", request.Repository.GetIntegration(),
	}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	var integration models.Integration
	err := res.Decode(&integration)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	// check GitHub integration connection and get last release
	client := github.NewClient(nil).WithAuthToken(integration.AccessToken)
	gitData := strings.Split(request.Repository.GetFullName(), "/")
	_, _, err = client.Repositories.GetLatestRelease(context.TODO(), gitData[0], gitData[1])
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	_, err = db.GetMongoClient().Database("registry").Collection("content").InsertOne(context.TODO(), model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusCreated, u.ContentResponse{
		Name:        request.Name,
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

	// get content
	collection := db.GetMongoClient().Database("registry").Collection("content")
	res := collection.FindOne(context.TODO(), bson.D{{"type", t}, {"name", name}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}
	var model models.Content
	err := res.Decode(&model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	// get integration
	res = db.GetMongoClient().Database("registry").Collection("integrations").FindOne(context.TODO(), bson.D{{
		"name", model.Repository.Integration,
	}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	var integration models.Integration
	err = res.Decode(&integration)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	// check GitHub integration connection
	client := github.NewClient(nil).WithAuthToken(integration.AccessToken)
	gitData := strings.Split(model.Repository.FullName, "/")
	_, _, err = client.Repositories.GetLatestRelease(context.TODO(), gitData[0], gitData[1])
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	model.UpdatedAt = time.Now().Format(time.RFC3339)
	model.Description = request.Description
	model.IsPublic = request.IsPublic

	update := bson.D{{"$set", bson.D{
		{"updated_at", model.UpdatedAt},
		{"description", model.Description},
		{"is_public", model.IsPublic},
	}}}
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"name", name}, {"type", t}}, update)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	cr, err := time.Parse(time.RFC3339, model.CreatedAt)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	up, err := time.Parse(time.RFC3339, model.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	c.JSON(http.StatusOK, u.ContentResponse{
		Name:        model.Name,
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

	res = db.GetMongoClient().Database("registry").Collection("content_revisions").FindOneAndDelete(context.TODO(), bson.D{
		{"content_name", request.Name},
	})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, u.DeleteContentResponse{Message: "Deleted"})
}

// GetDownloadLink works only for public content. For private content look at GetPluginContent and GetExtensionContent methods
func GetDownloadLink(c *gin.Context) {
	var request u.GetContentDownloadURIParams
	request.Name = c.Param("name")
	contentType := c.Param("type")
	request.Version = c.Param("version")

	// get content
	collection := db.GetMongoClient().Database("registry").Collection("content")
	res := collection.FindOne(context.TODO(), bson.D{{"type", contentType}, {"name", request.Name}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}
	var model models.Content
	err := res.Decode(&model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}
	if model.IsPublic == false {
		c.JSON(http.StatusForbidden, u.ForbiddenResponse{Message: "Attempt to download private content.", Code: http.StatusForbidden})
		return
	}

	// get integration
	res = db.GetMongoClient().Database("registry").Collection("integrations").FindOne(context.TODO(), bson.D{{
		"name", model.Repository.Integration,
	}})
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, u.ValidationErrorResponse{Message: res.Err().Error(), Code: http.StatusNotFound})
		return
	}

	var integration models.Integration
	err = res.Decode(&integration)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	contentPath := os.Getenv("CONTENT_DATA_FOLDER") + "/" + model.Name + "/" + request.Version + "/" + model.Name + ".zip"
	client := github.NewClient(nil).WithAuthToken(integration.AccessToken)
	gitData := strings.Split(model.Repository.FullName, "/")

	tags, _, err := client.Repositories.ListTags(context.TODO(), gitData[0], gitData[1], &github.ListOptions{})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
		return
	}

	// if file does not exist yet download it
	if _, err := os.Stat(contentPath); errors.Is(err, os.ErrNotExist) {
		for _, tag := range tags {
			if request.Version == tag.GetName() {
				err = downloadFile(contentPath, tag.GetZipballURL(), client.Client())
				if err != nil {
					c.JSON(http.StatusUnprocessableEntity, u.ValidationErrorResponse{Message: err.Error(), Code: http.StatusUnprocessableEntity})
					return
				}
				break
			}
		}
	}
	c.File(contentPath)
}

func downloadFile(path string, url string, client *http.Client) error {
	dir, _ := filepath.Split(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
