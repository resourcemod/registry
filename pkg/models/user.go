// Package models provide model structs
package models

type User struct {
	Name        string `bson:"name"`
	Password    string `bson:"password"`
	AccessToken string `bson:"access_token"`
	ExpiredAt   string `bson:"expired_at"`
	CreatedAt   string `bson:"created_at"`
	UpdatedAt   string `bson:"updated_at"`
}
