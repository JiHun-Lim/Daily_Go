package main

import "fmt"

func main() {

	type Product struct {
		name, category string
		price float64
	}

	var kayak = Product { "Kayak", "Watersports", 275.00}

	fmt.Println("Name is ", kayak.name, "\nCategory is ", kayak.category, "\nPrice is ", kayak.price)
}