package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func newProduct(name, category string, price float64) *Product {
	return &Product{ name, category, price}
}

func (product *Product) printDetails() {
	fmt.Println("Name : ", product.name, "Category : ", product.category, "Price : ", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + product.price * rate
	} else {
		return product.price
	}
}

func main() {

	products := [] * Product{
		{"Kayak", "Watersports", 275},
		{"LifeJacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for _, p := range products {
		p.printDetails()
	}


}