package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) { 
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

func main() {

// 	kayak := Product { "Kayak", "WaterSport", 275}
// 	insurance := Service { "Boat Cover", 12, 89.50}

// 	fmt.Println("Product : ", kayak.name, "Category : ", kayak.category, "Price : ", kayak.price)
// 	fmt.Println("Service : ", insurance.description, "Price : ", insurance.monthlyfee * float64(insurance.durationMonths))
// 
	expenses := []Expense {
		Product { "Kayak", "WaterSport", 275},
		Service { "Boat Cover", 12, 89.50},
	}

	for _, expense := range expenses {
		fmt.Println("Name : ", expense.getName(), "Cost : ", expense.getCost(true))
	}

	fmt.Println("Total : ", calcTotal(expenses))
}