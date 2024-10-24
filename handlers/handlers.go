package handlers

import (
	"html/template"
	"inventory-app/models"
	"net/http"
	"strconv"
)

// HomeHandler serves the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view", http.StatusSeeOther)
}

// AddProductHandler handles adding a new product.
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		category := r.FormValue("category")
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))              // Convert to int
		unitPrice, _ := strconv.ParseFloat(r.FormValue("unit_price"), 64) // Convert to float64

		models.AddProduct(name, category, quantity, unitPrice) // Call to add product in DB
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/add_product.html"))
	tmpl.Execute(w, nil)
}

// ViewProductsHandler displays the list of products.
func ViewProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts() // Example function to get products
	tmpl := template.Must(template.ParseFiles("templates/view_products.html"))
	tmpl.Execute(w, products)
}

// UpdateProductHandler handles updating an existing product.
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		category := r.FormValue("category")
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))              // Convert to int
		unitPrice, _ := strconv.ParseFloat(r.FormValue("unit_price"), 64) // Convert to float64

		models.UpdateProduct(id, name, category, quantity, unitPrice) // Call to update product in DB
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	// Logic to fetch the product details for updating
	id := r.URL.Query().Get("id")
	product := models.GetProductByID(id) // Fetch product by ID
	tmpl := template.Must(template.ParseFiles("templates/update_product.html"))
	tmpl.Execute(w, product)
}

// ReportHandler generates and serves reports.
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic to generate reports
}
