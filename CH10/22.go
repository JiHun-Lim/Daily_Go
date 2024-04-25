package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func newProduct(name, category string, price float64) *Product {
	// return &Product{name, category, price}
	return &Product{name, category, price - 10}
}

func main() {

	products := [2]*Product {
		newProduct("Kayak", "Watersports", 275.00),
		newProduct("Stadium", "Soccer", 75000),
	}

	for _, p := range products {
		fmt.Println(p.name, p.category, p.price)
	}
}