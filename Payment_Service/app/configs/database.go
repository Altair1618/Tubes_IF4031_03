package configs

import (
	"context"
	"log"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		uri := viper.Get("MONGODB_URI").(string)
		if uri == "" {
			log.Fatal("MONGODB_URI is not defined on .env")
		}

		username := viper.Get("MONGO_INITDB_ROOT_USERNAME").(string)
		password := viper.Get("MONGO_INITDB_ROOT_PASSWORD").(string)
		if username == "" || password == "" {
			log.Fatal("Incomplete DB credentials on .env")
		}
		clientOptions := options.Client().ApplyURI(uri).SetAuth(options.Credential{
			Username: username,
			Password: password,
		})

		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
