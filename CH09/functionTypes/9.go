package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(prodcuct string, price float64, calculator calcFunc) {
	fmt.Println("Product :", prodcuct, "Price: ", calculator(price))
}

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		var withtax calcFunc = func (price float64) float64 {
			return price + price * 0.2
		}
		return withtax
		} else {
		var withouttax calcFunc = func (price float64) float64 {
			return price
		}
		return withouttax
		// return withtax
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