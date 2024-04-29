package main

import (
	"fmt"
	"packages/store"
)

func main() {
	product := store.NewProduct("Kayak", "WaterSports", 275)

	product.SetPrice(8000)

	fmt.Println(product.Name, product.Category, product.Price())

}