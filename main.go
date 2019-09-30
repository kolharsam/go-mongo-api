package main

import (
	"some-rest-api/logger"
	"some-rest-api/mongo"
	"some-rest-api/server"
)

func init() {
	mongo.GoMongo()
}

func main() {
	if mongo.IsLive() {
		server.FireServer()
	} else {
		logger.Message("The API couldn't be setup.", "fatal")
	}
}
