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

func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithtax
	} else {
		return calcWithouttax
	}
}

func main(){
	products := map[string]float64 {
		"Kayak" : 275,
		"LifeJacket" : 48.95,
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}