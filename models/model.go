package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// Product represents a product in the inventory.
type Product struct {
	ID        int
	Name      string
	Category  string
	Quantity  int
	UnitPrice float64
}

// DB is the database connection
var DB *sql.DB

// Initialize the database
func init() {
	var err error
	DB, err = sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		category TEXT,
		quantity INTEGER,
		unit_price REAL
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// AddProduct adds a new product to the database.
func AddProduct(name, category string, quantity int, unitPrice float64) {
	query := `INSERT INTO products (name, category, quantity, unit_price) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, name, category, quantity, unitPrice)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllProducts retrieves all products from the database.
func GetAllProducts() []Product {
	query := `SELECT id, name, category, quantity, unit_price FROM products`
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Quantity, &product.UnitPrice); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products
}

// GetProductByID retrieves a product by its ID.
func GetProductByID(id string) Product {
	query := `SELECT id, name, category, quantity, unit_price FROM products WHERE id = ?`
	row := DB.QueryRow(query, id)

	var product Product
	if err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Quantity, &product.UnitPrice); err != nil {
		log.Fatal(err)
	}
	return product
}

// UpdateProduct updates an existing product in the database.
func UpdateProduct(id, name, category string, quantity int, unitPrice float64) {
	query := `UPDATE products SET name = ?, category = ?, quantity = ?, unit_price = ? WHERE id = ?`
	_, err := DB.Exec(query, name, category, quantity, unitPrice, id)
	if err != nil {
		log.Fatal(err)
	}
}
