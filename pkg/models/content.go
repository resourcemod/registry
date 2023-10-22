// Package models provide model structs
package models

const (
	TYPE_PLUGIN    string = "content"
	TYPE_EXTENSION string = "extension"
)

type Content struct {
	Name        string `bson:"name"`
	Version     string `bson:"version"`
	Type        string `bson:"type"`
	Description string `bson:"description"`
	IsPublic    bool   `bson:"is_public"`
	UserName    string `bson:"user_name"`
	Repository  string `bson:"repository"`
	CreatedAt   string `bson:"created_at"`
	UpdatedAt   string `bson:"updated_at"`
}
