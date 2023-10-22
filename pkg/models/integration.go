// Package models provide model structs
package models

const (
	TYPE_GIT    string = "git"
	HOST_GITHUB string = "github.com"
)

type Integration struct {
	Type        string `bson:"type"`
	Name        string `bson:"name"`
	Host        string `bson:"host"`
	AccessToken string `bson:"access_token"`
}
