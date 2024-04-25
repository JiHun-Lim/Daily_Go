package main

import "fmt"

type Product struct {
	name, category string
	price float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price - 10, supplier}
}

func copyProduct(product *Product) Product {
	p := *product
	s:= *product.Supplier
	p.Supplier = &s
	return p
}

func main() {

	acme := &Supplier {"Acme Co", "New York"}
	
	p1 := newProduct("Kayak", "Watersports", 275, acme)

	p2 := copyProduct(p1)

	p1.name = "Original kayak"

	p2.Supplier.name = "BoatCo"

	for _, p := range[]Product {*p1, p2} {
		fmt.Println(p.name, p.Supplier.name, p.Supplier.city)
	}

}