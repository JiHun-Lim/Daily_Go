package main

import "fmt"

func main() {

	type Product struct {
		name, category string
		price float64
	}

	type StockLevel struct {
		Product
		Alternate Product
		count int
	}

	array := [1]StockLevel{
		{
			Product : Product{ "Kayak", "Watersports", 275.00},
			Alternate : Product{ "Stadium", "Soccer", 75000},
			count : 100,
		},
	}

	fmt.Println("Array:", array[0].Product.name)
}