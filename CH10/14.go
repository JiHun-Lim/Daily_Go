package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

func main() {

	type Product struct {
		name, category string
		price float64
	}

	prod := Product { "Kayak", "Watersports", 275.00};

	var builder strings.Builder

	json.NewEncoder(&builder).Encode(struct{
		ProductName string
		ProductPrice float64
	}{
		ProductName: prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println(builder.String())
}
