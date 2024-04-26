package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func main() {

	products := [] * Product{
		{"Kayak", "Watersports", 275},
		{"LifeJacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for _, p := range products {

		fmt.Println("Name : ", p.name, "Category : ", p.category, "Price : ", p.price)
	}
}