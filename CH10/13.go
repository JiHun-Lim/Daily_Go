package main

import "fmt"

func writeName (val struct{
	name, category string
	price float64}) {
	fmt.Println(val.name)
}

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
	item := Item { "Stadium", "Soccer", 75000};

	writeName(prod)
	writeName(item)

}
