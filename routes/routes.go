package routes

import (
	"github.com/AdrianRamsey13/backend-ecommerce/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Register the routes
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
}
