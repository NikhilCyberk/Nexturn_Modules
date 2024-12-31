/*
Exercise 3: Inventory Management System

Topics Covered: Go Conditions, Go Type Casting, Go Functions, Go Arrays, Go Strings, Go Errors

Case Study:

A store needs to manage its inventory of products. Build an application that includes the following:

1. Product Struct: Create a struct to represent a product with fields for ID, name, price (float64), and stock (int).

2. Add Product: Write a function to add new products to the inventory. Use type casting to ensure price inputs are converted to float64.

3. Update Stock: Implement a function to update the stock of a product. Use conditions to validate the input (e.g., stock cannot be negative).

4. Search Product: Allow users to search for products by name or ID. If a product is not found, return a custom error message.

5. Display Inventory: Use loops to display all available products in a formatted table.

Bonus:

• Add sorting functionality to display products by price or stock in ascending order.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Constants for validation
const (
	MIN_PRICE     = 0.01
	MAX_PRICE     = 999999.99
	MAX_STOCK     = 999999
	LOW_STOCK     = 10
	CRITICAL_STOCK = 5
)

// Custom error type for inventory operations
type InventoryError struct {
	Field   string
	Message string
}

func (e *InventoryError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type Product struct {
	ID        int
	Name      string
	Price     float64
	Stock     int
	Status    string
	Category  string
	UpdatedAt time.Time
}

// validateProduct checks if product data is valid
func validateProduct(p Product, inventory []Product) error {
	// Validate ID
	if p.ID <= 0 {
		return &InventoryError{Field: "ID", Message: "must be positive"}
	}
	for _, product := range inventory {
		if product.ID == p.ID {
			return &InventoryError{Field: "ID", Message: "already exists"}
		}
	}

	// Validate Name
	if strings.TrimSpace(p.Name) == "" {
		return &InventoryError{Field: "Name", Message: "cannot be empty"}
	}

	// Validate Price
	if p.Price < MIN_PRICE || p.Price > MAX_PRICE {
		return &InventoryError{
			Field:   "Price",
			Message: fmt.Sprintf("must be between %.2f and %.2f", MIN_PRICE, MAX_PRICE),
		}
	}

	// Validate Stock
	if p.Stock < 0 || p.Stock > MAX_STOCK {
		return &InventoryError{
			Field:   "Stock",
			Message: fmt.Sprintf("must be between 0 and %d", MAX_STOCK),
		}
	}

	// Validate Category
	if strings.TrimSpace(p.Category) == "" {
		return &InventoryError{Field: "Category", Message: "cannot be empty"}
	}

	return nil
}

// readInput reads and validates user input
func readInput(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	return strings.TrimSpace(scanner.Text()), nil
}

// updateProductStatus updates the status based on stock level
func updateProductStatus(stock int) string {
	switch {
	case stock == 0:
		return "Out of Stock"
	case stock <= CRITICAL_STOCK:
		return "Critical Stock"
	case stock <= LOW_STOCK:
		return "Low Stock"
	default:
		return "In Stock"
	}
}

func addProduct(inventory *[]Product) error {
	fmt.Println("\n=== Add New Product ===")

	// Read ID
	idStr, err := readInput("Enter product ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &InventoryError{Field: "ID", Message: "must be a number"}
	}

	// Read Name
	name, err := readInput("Enter product name: ")
	if err != nil {
		return err
	}

	// Read Price
	priceStr, err := readInput("Enter product price: ")
	if err != nil {
		return err
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return &InventoryError{Field: "Price", Message: "must be a number"}
	}

	// Read Stock
	stockStr, err := readInput("Enter product stock: ")
	if err != nil {
		return err
	}
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return &InventoryError{Field: "Stock", Message: "must be a number"}
	}

	// Read Category
	category, err := readInput("Enter product category: ")
	if err != nil {
		return err
	}

	newProduct := Product{
		ID:        id,
		Name:      name,
		Price:     price,
		Stock:     stock,
		Category:  category,
		Status:    updateProductStatus(stock),
		UpdatedAt: time.Now(),
	}

	if err := validateProduct(newProduct, *inventory); err != nil {
		return err
	}

	*inventory = append(*inventory, newProduct)
	fmt.Printf("\n✅ Product '%s' added successfully!\n", name)
	return nil
}

func updateStock(inventory *[]Product) error {
	fmt.Println("\n=== Update Product Stock ===")

	// Read ID
	idStr, err := readInput("Enter product ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &InventoryError{Field: "ID", Message: "must be a number"}
	}

	// Find product
	var found bool
	for i, product := range *inventory {
		if product.ID == id {
			stockStr, err := readInput("Enter new stock quantity: ")
			if err != nil {
				return err
			}
			stock, err := strconv.Atoi(stockStr)
			if err != nil {
				return &InventoryError{Field: "Stock", Message: "must be a number"}
			}

			if stock < 0 || stock > MAX_STOCK {
				return &InventoryError{
					Field:   "Stock",
					Message: fmt.Sprintf("must be between 0 and %d", MAX_STOCK),
				}
			}

			(*inventory)[i].Stock = stock
			(*inventory)[i].Status = updateProductStatus(stock)
			(*inventory)[i].UpdatedAt = time.Now()
			
			fmt.Printf("\n✅ Stock updated for '%s'\n", product.Name)
			if stock <= LOW_STOCK {
				fmt.Printf("⚠️ Warning: Low stock alert (%d items remaining)\n", stock)
			}
			found = true
			break
		}
	}

	if !found {
		return &InventoryError{Field: "Product", Message: "not found"}
	}
	return nil
}

func searchProduct(inventory []Product) error {
	fmt.Println("\n=== Search Product ===")
	
	query, err := readInput("Enter product ID or name to search: ")
	if err != nil {
		return err
	}

	found := false
	fmt.Println("\nSearch Results:")
	fmt.Println(strings.Repeat("-", 80))

	for _, p := range inventory {
		// Try to convert query to ID for comparison
		queryID, err := strconv.Atoi(query)
		if (err == nil && p.ID == queryID) || strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) {
			displayProduct(p)
			found = true
		}
	}

	if !found {
		return &InventoryError{Field: "Search", Message: "no products found"}
	}
	return nil
}

func displayProduct(p Product) {
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Category: %s\n", p.Category)
	fmt.Printf("Price: $%.2f\n", p.Price)
	fmt.Printf("Stock: %d\n", p.Stock)
	fmt.Printf("Status: %s\n", p.Status)
	fmt.Printf("Last Updated: %s\n", p.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println(strings.Repeat("-", 80))
}

func displayInventory(inventory []Product) {
	if len(inventory) == 0 {
		fmt.Println("\n❌ No products in inventory.")
		return
	}

	fmt.Println("\n=== Current Inventory ===")
	fmt.Println(strings.Repeat("=", 130))
	fmt.Printf("%-5s | %-20s | %-20s | %-10s | %-8s | %-15s | %-10s\n",
		"ID", "Name", "Category", "Price", "Stock", "Status", "Updated")
	fmt.Println(strings.Repeat("-", 130))

	for _, p := range inventory {
		fmt.Printf("%-5d | %-20s | %-20s | $%-9.2f | %-8d | %-15s | %s\n",
			p.ID,
			truncateString(p.Name, 20),
			truncateString(p.Category, 20),
			p.Price,
			p.Stock,
			p.Status,
			p.UpdatedAt.Format("2006-01-02"))
	}
	fmt.Println(strings.Repeat("=", 130))
}

func truncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}

func sortInventory(inventory *[]Product) error {
	fmt.Println("\n=== Sort Inventory ===")
	fmt.Println("1. Sort by Price (ascending)")
	fmt.Println("2. Sort by Price (descending)")
	fmt.Println("3. Sort by Stock (ascending)")
	fmt.Println("4. Sort by Stock (descending)")
	fmt.Println("5. Sort by Name")
	
	choice, err := readInput("\nEnter your choice (1-5): ")
	if err != nil {
		return err
	}

	switch choice {
	case "1":
		sort.Slice(*inventory, func(i, j int) bool {
			return (*inventory)[i].Price < (*inventory)[j].Price
		})
	case "2":
		sort.Slice(*inventory, func(i, j int) bool {
			return (*inventory)[i].Price > (*inventory)[j].Price
		})
	case "3":
		sort.Slice(*inventory, func(i, j int) bool {
			return (*inventory)[i].Stock < (*inventory)[j].Stock
		})
	case "4":
		sort.Slice(*inventory, func(i, j int) bool {
			return (*inventory)[i].Stock > (*inventory)[j].Stock
		})
	case "5":
		sort.Slice(*inventory, func(i, j int) bool {
			return (*inventory)[i].Name < (*inventory)[j].Name
		})
	default:
		return &InventoryError{Field: "Sort", Message: "invalid choice"}
	}

	fmt.Println("\n✅ Inventory sorted successfully!")
	displayInventory(*inventory)
	return nil
}



func main() {
	inventory := []Product{
		{
			ID:        1,
			Name:      "Laptop",
			Price:     999.99,
			Stock:     15,
			Category:  "Electronics",
			Status:    "In Stock",
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Smartphone",
			Price:     499.99,
			Stock:     8,
			Category:  "Electronics",
			Status:    "Low Stock",
			UpdatedAt: time.Now(),
		},
	}

	menu := `
=== Inventory Management System ===
1. Add Product
2. Update Stock
3. Search Product
4. Display Inventory
5. Sort Inventory
6. Exit

Enter your choice (1-6): `

	for {
		choice, err := readInput(menu)
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		var operationErr error
		switch choice {
		case "1":
			operationErr = addProduct(&inventory)
		case "2":
			operationErr = updateStock(&inventory)
		case "3":
			operationErr = searchProduct(inventory)
		case "4":
			displayInventory(inventory)
		case "5":
			operationErr = sortInventory(&inventory)
		case "6":
			fmt.Println("\n👋 Thank you for using the Inventory Management System. Goodbye!")
			return
		default:
			fmt.Println("\n❌ Invalid choice. Please try again.")
			continue
		}

		if operationErr != nil {
			fmt.Printf("\n❌ Error: %v\n", operationErr)
		}
	}
}