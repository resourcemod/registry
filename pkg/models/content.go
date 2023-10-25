// Package models provide model structs
package models

const (
	TYPE_PLUGIN    string = "plugin"
	TYPE_EXTENSION string = "extension"
)

type Repository struct {
	Integration string `bson:"integration"`
	FullName    string `bson:"full_name"`
	GitUrl      string `bson:"git_url"`
}

type Content struct {
	Name        string     `bson:"name"`
	Type        string     `bson:"type"`
	Description string     `bson:"description"`
	IsPublic    bool       `bson:"is_public"`
	UserName    string     `bson:"user_name"`
	Repository  Repository `bson:"repository"`
	CreatedAt   string     `bson:"created_at"`
	UpdatedAt   string     `bson:"updated_at"`
}
