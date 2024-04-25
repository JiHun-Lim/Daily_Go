package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func calcTax(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}

func main() {


	kayak := &Product {
		name : "Kayak",
		category : "Watersports",
		price : 275,
	}

	fmt.Println("Before : ", kayak.name, kayak.category, kayak.price)

	calcTax(kayak)
	
	fmt.Println("After : ", kayak.name, kayak.category, kayak.price)
}