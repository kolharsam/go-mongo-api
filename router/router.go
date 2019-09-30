package router

import (
	"some-rest-api/routes"

	"github.com/gorilla/mux"
)

var appRouter *mux.Router

func init() {
	appRouter = mux.NewRouter().StrictSlash(true)
}

// GetRouter := Get app Router that was setup
func GetRouter() *mux.Router {
	return appRouter
}

// SetupRoutes := Get all the routes set up for the router
func SetupRoutes() {
	appRouter.HandleFunc("/", routes.Home).Methods("GET")
	appRouter.HandleFunc("/movie", routes.AddMovie).Methods("POST")
	appRouter.HandleFunc("/movie/{id}", routes.GetMovie).Methods("GET")
}
