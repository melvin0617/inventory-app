package main

import (
	"inventory-app/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/add", handlers.AddProductHandler)
	http.HandleFunc("/view", handlers.ViewProductsHandler)
	http.HandleFunc("/update", handlers.UpdateProductHandler)
	http.HandleFunc("/report", handlers.ReportHandler)

	http.ListenAndServe(":8080", nil)
}
