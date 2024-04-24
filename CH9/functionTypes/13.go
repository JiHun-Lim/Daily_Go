package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(prodcuct string, price float64, calculator calcFunc) {
	fmt.Println("Product :", prodcuct, "Price: ", calculator(price))
}

func main(){
	watersportsProducts := map[string]float64 {
		"Kayak" : 275,
		"Lifejacket" : 48.95,
	}

	soccerProducts := map[string]float64 {
		"SoccerBall" : 19.50,
		"Stadium" : 79500,
	}

	calc := func(price float64)	float64 {
		if price > 100 {
			return price * 1.2
		} else{
			return price;
		}
	}

	for product,price := range watersportsProducts {
		printPrice(product,price, calc)
	}

	calc = func(price float64) float64 {
		if price > 50 {
			return price * 1.1
		} else {
			return price
		}
	}

	for product,price := range soccerProducts {
		printPrice(product,price, calc)
	}


}