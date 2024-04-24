package main

import "fmt"

type calcFunc func(float64) float64

var prizeGiveaway = false

func printPrice(prodcuct string, price float64, calculator calcFunc) {
	fmt.Println("Product :", prodcuct, "Price: ", calculator(price))
}

func priceCalc(threshold float64, rate float64, zeroPrices *bool) calcFunc {
	return func(price float64) float64 {
		if *zeroPrices {
			return 0
		} else{
			if price > threshold {
				return price * (1 + rate)
			} else{
				return price
			}	

		}
	}
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

	prizeGiveaway = false
	watercalc := priceCalc(100, 0.2, &prizeGiveaway);
	prizeGiveaway = true
	soccercalc := priceCalc(50, 0.1, &prizeGiveaway);

	for product,price := range watersportsProducts {
		printPrice(product,price, watercalc)
	}

	for product,price := range soccerProducts {
		printPrice(product,price, soccercalc)
	}


}