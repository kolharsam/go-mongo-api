package router

import (
	"net/http"
	"some-rest-api/middleware"
	"some-rest-api/routes"

	"github.com/gorilla/mux"
)

type Route struct {
	route        string
	routeHandler func(w http.ResponseWriter, r *http.Request)
	method       string
}

var appRouter *mux.Router

func init() {
	appRouter = mux.NewRouter().StrictSlash(true)
	appRouter.Use(middleware.RequestLogger)
	appRouter.Use(middleware.AddHeaders)
}

// GetRouter := Get app Router that was setup
func GetRouter() *mux.Router {
	return appRouter
}

// SetupRoutes := Get all the routes set up for the router
func SetupRoutes() {
	appRouter.HandleFunc("/", routes.Home).Methods("GET")
	appRouter.HandleFunc("/movie", routes.Add).Methods("POST")
	appRouter.HandleFunc("/movie/{id}", routes.GetOne).Methods("GET")
	appRouter.HandleFunc("/movie", routes.Update).Methods("PUT")
	appRouter.HandleFunc("/movie/{id}", routes.Delete).Methods("DELETE")
	appRouter.HandleFunc("/movies", routes.GetAll).Methods("GET")
}
