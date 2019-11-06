package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "mongodb://mongodb0.example.com:27017/goweb")
