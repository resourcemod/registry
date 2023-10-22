// Package db handle database connection
package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

func InitDBConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	client = c

	db := c.Database("registry")
	command := bson.D{{"create", "users"}}
	var result bson.M
	if err = db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		log.Fatal(err)
	}
	command = bson.D{{"create", "content"}}
	if err = db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		log.Fatal(err)
	}
	command = bson.D{{"create", "integration"}}
	if err = db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		log.Fatal(err)
	}

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = db.Collection("users").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	_, err = db.Collection("content").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	_, err = db.Collection("integration").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
}

func GetMongoClient() *mongo.Client {
	if client == nil {
		InitDBConnection()
	}
	return client
}
