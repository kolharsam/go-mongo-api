package mongo

import (
	"context"
	"net/http"
	"some-rest-api/env"
	"some-rest-api/funcerr"
	"some-rest-api/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type envVars struct {
	mongoURL            string
	mongoDBName         string
	mongoCollectionName string
}

type mongoConfig struct {
	mongoClient     *mongo.Client
	mongoCollection *mongo.Collection
}

var allEnvVars envVars
var config mongoConfig
var isLive bool

func init() {
	allEnvVars.mongoURL = env.GetEnvVar("MONGO_URL")
	allEnvVars.mongoDBName = env.GetEnvVar("MONGO_DB_NAME")
	allEnvVars.mongoCollectionName = env.GetEnvVar("MONGO_COLLECTION")
}

func isConnected() bool {
	connect := false
	resp, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		logger.Message("Check your internet connection", "fatal")
	} else {
		switch resp.StatusCode {
		case 204:
			connect = true
		case 504:
			logger.Message("Server Error", "fatal")
		default:
			connect = true
		}
	}
	return connect
}

func setClient(client *mongo.Client) {
	config.mongoClient = client
}

func setCollection(collection *mongo.Collection) {
	config.mongoCollection = collection
}

func setupMongoCollection(client *mongo.Client, dbname string, collecName string) *mongo.Collection {
	return client.Database(dbname).Collection(collecName)
}

func connectMongo() bool {
	result := false
	if isConnected() {
		logger.Message("You're connected to the internet!", "info")
		client, err := mongo.NewClient(options.Client().ApplyURI(allEnvVars.mongoURL))
		if !funcerr.IsErrored(err, "There was an error connecting to Mongo!") {
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			newErr := client.Connect(ctx)
			if !funcerr.IsErrored(newErr, "Error occurred while setting up connection") {
				setClient(client)
				setCollection(
					setupMongoCollection(
						client,
						allEnvVars.mongoDBName,
						allEnvVars.mongoCollectionName,
					),
				)
				result = true
			}
		}
	}
	isLive = result
	return result
}

// GetClient := Getter for mongo client
func GetClient() *mongo.Client {
	return config.mongoClient
}

// GetCollection := Getter for mongo collection
func GetCollection() *mongo.Collection {
	return config.mongoCollection
}

// IsLive := Return value of isLive
func IsLive() bool {
	return isLive
}

// GoMongo := Connect to mongo with this.
func GoMongo() {
	if connectMongo() {
		logger.Message("Successfully connected to Mongo", "info")
	} else {
		logger.Message("There was a problem", "fatal")
	}
}
