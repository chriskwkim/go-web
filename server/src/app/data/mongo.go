package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var Client, err = mongo.Connect(context.Background(), "mongodb://localhost:27017/goweb", nil)
// Set client options
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB
var Client, err = mongo.Connect(context.TODO(), clientOptions)
