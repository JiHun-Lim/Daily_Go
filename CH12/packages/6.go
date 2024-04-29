package main

import (
	"fmt"
	"packages/store"
)

func main() {
	product := store.Product{
		Name : "Kayak",
		Category : "Watersports",
		price : 275,
	}

	fmt.Println(product.Name, product.Category)

}