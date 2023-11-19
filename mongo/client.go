package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func getCloudUrl() string {
	return os.Getenv("X-SongUser-MongoCloud-Url")
}

func getMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() { // Thread safe
		if clientInstance == nil { // Ensure the instance is initialized only once
			clientOptions := options.Client().ApplyURI(getCloudUrl())
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				clientInstanceError = err
			}
			clientInstance = client
		}
	})

	err := clientInstance.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return clientInstance, clientInstanceError
}
