package main

import "fmt"

func main() {

	type Product struct {
		name, category string
		price float64
	}

	type StockLevel struct {
		Product
		Alternate Product
		count int

	}

	stockItem := StockLevel {
		Product : Product { "Kayak", "Watersports", 275.00},
		Alternate : Product { "Lifejacket", "Watersports", 48.95},
		count : 100,
	}

	fmt.Println(stockItem.Product.name, stockItem.count)
	fmt.Println(stockItem.Alternate.name)
}