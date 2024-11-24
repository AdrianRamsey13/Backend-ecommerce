package controllers

import (
	"encoding/json"
	"net/http"
)

// GetProducts returns a list of products
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// GetProducts returns a list of products
var products = []Product{
	{ID: "1", Name: "Laptop", Price: 1000},
	{ID: "2", Name: "Mouse", Price: 20},
}

// GetProducts returns a list of products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Set the header
	w.Header().Set("Content-Type", "application/json")
	// Encode the products to JSON
	json.NewEncoder(w).Encode(products)
}
