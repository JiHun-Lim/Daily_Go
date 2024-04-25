package main

import "fmt"

func main() {

	type Product struct {
		name, category string
		price float64
	}

	type Item struct {
		name, category string
		price float64
	}

	prod := Product { "Kayak", "Watersports", 275.00};
	item := Item { "Kayak", "Watersports", 275.00};
	
	fmt.Println(prod == Product(item))
}
