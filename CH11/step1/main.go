package main

import "fmt"

func main() {

	kayak := Product { "Kayak", "WaterSport", 275}
	insurance := Service { "Boat Cover", 12, 89.50}

	fmt.Println("Product : ", kayak.name, "Category : ", kayak.category, "Price : ", kayak.price)
	fmt.Println("Service : ", insurance.description, "Price : ", insurance.monthlyfee * float64(insurance.durationMonths))
}