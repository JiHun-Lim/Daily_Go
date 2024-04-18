package main

import (
	"fmt"
	// "sort"
)

func main() {
	// fmt.Println("Hello Collections")

	// var names [3]string

	// names[0] = "kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"

	// fmt.Println(names)

	// names := [3]string{"kayak", "Lifejacket", "Paddle"}

	// names := [2][2]string{{"kayak", "Lifejacket"}, {"Paddle", "Jihun"}}

	// var otherArray []string = names
	// var otherArray [4]string = names

	// otherArray := names

	// anotherArray := &names

	// names[0] = "Canoe"

	// fmt.Println(names)
	// fmt.Println(otherArray)
	// fmt.Println(anotherArray, *anotherArray)

	// for index, value := range names {
	// 	fmt.Println(index, value)
	// }

	// names := make([]string, 30)

	// names := []string{"kayak", "Lifejacket", "Paddle"}

	// names = append(names, "Jihun")

	// names := make([]string, 3,6)

	// names = append(names, "kayak", "Lifejacket", "Paddle", "kayak", "Lifejacket", "Paddle")

	// names := make([]string, 3,6)

	// names[0] = "Jihun"

	// othernames := []string {"kayak", "Lifejacket", "Paddle"}

	// newnames := append(names, othernames ...)

	// fmt.Println(newnames)

	// fmt.Println(newnames[2:5], len(newnames[2:5]), cap(newnames))

	// products := [4]string {"kayak", "Lifejacket", "Paddle", "hat"}

	// allNames := products[1:]
	// someNames := allNames[1:3]

	// allNames = append(allNames, "gloves")
	// allNames[1] = "Canoe"

	// fmt.Println("someNames : ", someNames)
	// fmt.Println("allNames : ", allNames)

	// products := [4]string {"kayak", "Lifejacket", "Paddle", "hat"}

	// allNames := products[1:]
	// // someNames := make([]string, 2)
	// var someNames []string
	// copy(someNames, allNames)

	// fmt.Println(someNames, allNames)

	// products := []string {"kayak", "Lifejacket", "Paddle", "hat"}

	// sort.Strings(products)

	// fmt.Println(products)

	// products := map[string]float64 {
	// 	"kayak" : 279,
	// 	"lifejacket" : 48.95,
	// 	"Hat" : 0,
	// }

	// fmt.Println(products, products["Hat"])

	var price = "$48.95"

	for index, char := range price{
		fmt.Println(index, char, string(char))
	}

}