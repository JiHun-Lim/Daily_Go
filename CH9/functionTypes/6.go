package main

import "fmt"

func calcWithtax(price float64) float64 {
	return price + price * 0.2
}

func calcWithouttax(price float64) float64 {
	return price
}

func printPrice(prodcuct string, price float64, calculator func(float64) float64) {
	fmt.Println("Product :", prodcuct, "Price: ", calculator(price))
}

func main(){
	products := map[string]float64 {
		"Kayak" : 275,
		"LifeJacket" : 48.95,
	}

	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithtax)
		} else{
			printPrice(product, price, calcWithouttax)
		}

	}
}