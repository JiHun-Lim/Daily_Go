package main

import (
	"fmt"
	// "strconv"
)

func main() {

	// kayakPrice := 275.00
	// fmt.Println("Price :", kayakPrice)

	// if kayakPrice > 100{
	// 	fmt.Println("Price is greater than 100")
	// }

	// if kayakPrice > 500 {
	// 	fmt.Println("Price is greater than 500")
	// } else if kayakPrice < 300 {
	// 	fmt.Println("Price is lower than 300")
	// }

	// if kayakPrice > 500 {
	// 	scopedVar := 500
	// 	fmt.Println("Price is greater than", scopedVar)
	// } else if kayakPrice < 100 {
	// 	scopedVar := "Price is less than 100"
	// 	fmt.Println(scopedVar)
	// } else {
	// 	scopedVar := false
	// 	fmt.Println("Matched: ", scopedVar)
	// }

	// priceString := "275"
	// if kayakPrice, err := strconv.Atoi(priceString); err == nil{
	// 	fmt.Println("Price: ", kayakPrice)
	// } else {
	// 	fmt.Println("Error: ", err)
	// }

	// counter := 0

	// for {
	// 	fmt.Println("Counter :", counter)
	// 	counter++
	// 	if counter >3 {
	// 		break
	// 	}
	// }

	// for counter <= 3 {
	// 	fmt.Println("Counter :", counter)
	// 	counter++
	// }

	// for counter :=0; counter <=3; counter++ {
	// 	fmt.Println("Counter :", counter)
	// }

	// for counter :=0; true; counter++ {
	// 	fmt.Println("Counter :", counter)
	// 	if counter > 3{
	// 		break
	// 	} else {
	// 		continue
	// 	}
	// }

	// product := "Kayak"

	// for index, character := range product {
	// 	fmt.Println("Index :", index, "Character :", string(character))
	// }

	// product := "Kayak"

	// for _, character := range product {
	// 	fmt.Println("Character :", string(character))
	// }

	// products := []string {"Kayak", "Lifejacket", "Soccer Ball"}

	// for index, element := range(products){
	// 	fmt.Println("Index :", index, "Element :", element)
	// }

	// product := "Kayak"

	// for index, character := range product {
	// 	switch character {
	// 	case 'K' :
	// 		fmt.Println("K at position", index)
	// 	case 'y' :
	// 		fmt.Println("y at position", index)
	// 	}

	// }

	// for index, character := range product {
	// 	switch character {
	// 	case 'K' :
	// 		fmt.Println("Uppercase character")
	// 		fallthrough
	// 	case 'k':
	// 		fmt.Println("k at position", index)
	// 		// fallthrough
	// 	case 'y' :
	// 		fmt.Println("y at position", index)
	// 	default:
	// 		fmt.Println("Character", string(character), "at position", index)
	// 	}
	
	counter := 0
	target : fmt.Println("Counter", counter)
	counter++

	if counter < 5{
		goto target
	}

}