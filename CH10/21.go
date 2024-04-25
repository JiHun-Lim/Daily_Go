package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

func main() {


	kayak := calcTax(&Product {
		name : "Kayak",
		category : "Watersports",
		price : 275,
	})
	
	fmt.Println(kayak.name, kayak.category, kayak.price)
}