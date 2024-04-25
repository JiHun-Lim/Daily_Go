package main

import "fmt"

func main() {

	type Product struct {
		name, category string
		price float64
	}

	type StockLevel struct {
		Product
		count int

	}

	stockItem := StockLevel {
		Product : Product { "Kayak", "Watersports", 275.00},
		count : 100,
	}

	fmt.Println(stockItem.Product.name, stockItem.count)
}