package main

import (
	"log"
	"net/http"

	"github.com/AdrianRamsey13/backend-ecommerce/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Register the routes
	routes.RegisterRoutes(router)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
