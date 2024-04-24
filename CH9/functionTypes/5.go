package main

import "fmt"

func calcWithtax(price float64) float64 {
	return price + price * 0.2
}

func calcWithouttax(price float64) float64 {
	return price
}

func main(){
	products := map[string]float64 {
		"Kayak" : 275,
		"LifeJacket" : 48.95,
	}

	for product, price := range products {
		var calcFunc func(float64) float64

		fmt.Println("Function assigned: ", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithtax
		} else{
			calcFunc = calcWithouttax
		}

		fmt.Println("Function assigned: ", calcFunc == nil)

		totalPrice := calcFunc(price)
		fmt.Println("Product :", product, "Price :", price, "Total Price", totalPrice)
	}
}