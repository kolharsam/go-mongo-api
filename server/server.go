package server

import (
	"log"
	"net/http"
	"some-rest-api/env"
	"some-rest-api/router"
)

// Start := to start server
func Start() {
	port := env.GetEnvVar("SERVER_PORT")
	appRouter := router.GetRouter()
	router.SetupRoutes()
	err := http.ListenAndServe(port, appRouter)
	log.Fatal(err)
}
