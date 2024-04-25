package main

import "fmt"

func main() {

	type Product struct {
		name, category string
		price float64
	}

	kayak := Product {
		name : "Kayak",
		category : "Watersports",
	}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300

	fmt.Println("Price changed to ", kayak.price)

	var lifejacket Product

	fmt.Println("For Lifejacket, Name is ", lifejacket.name, "Category is ", lifejacket.category, "Price is ", lifejacket.price)
}