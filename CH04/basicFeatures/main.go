package main

import (
	"fmt"
	"sort"
	// "math/rand"
)

func main(){
	// fmt.Println("Value :", rand.Int())
	
	// fmt.Println("Hello, Go")
	// fmt.Println(20 + 20)
	// fmt.Println(20 + 30)

	// const price float32 = 275.00
	// const tax float32 = 27.50
	// fmt.Println(price + tax)
	// const quantity = 2
	// fmt.Println("Total:", quantity * (price + tax))

	// const price, tax float32 = 275, 277.50
	// const quantity, inStock = 2, true
	// fmt.Println("Total:", quantity * (price + tax))
	// fmt.Println("In stock:", inStock)

	// var price float32 = 275.00
	// var tax float32 = 27.50
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)

	// var price = 275.00
	// var price2 = price
	// var tax float32 = 27.50
	// fmt.Println(price)
	// fmt.Println(price2)
	// fmt.Println(price + tax)

	// var price float32
	// fmt.Println(price)
	// price = 275.00
	// fmt.Println(price)

	// var price, tax = 275.00, 27.50
	// fmt.Println(price + tax)

	// var price, tax float64
	// price = 275.00
	// tax = 27.50
	// fmt.Println(price + tax)

	// price := 275.00
	// fmt.Println(price)

	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price + tax)
	// fmt.Println("In stock:", inStock)

	// var price2, tax = 200.00, 25.00
	// price2, tax, _ := 200.00, 25.00, true

	// var _ = "Alice"

	// fmt.Println("Toatal 2:", price2 + tax)

	// first := 100
	// second := first

	// first++

	// fmt.Println(first)
	// fmt.Println(second)

	// first := 100
	// var second *int = &first

	// first++
	// *second++

	// var mypointer *int

	// mypointer = second
	// *mypointer++

	// fmt.Println(first)
	// fmt.Println(*second)

	// first := 100
	// var second *int

	// fmt.Println(second)
	// second = &first
	// fmt.Println(second)
	// fmt.Println(second == nil)

	// first := 100
	// second := &first
	// third := &second

	// fmt.Println(first)
	// fmt.Println(*second)
	// fmt.Println(**third)

	names := [3]string {"Alice", "Charlie", "Bob"}

	// secondPosition := names[1]

	// fmt.Println(secondPosition)

	// sort.Strings(names[:])

	// fmt.Println(secondPosition)

	secondPosition := &names[1]

	fmt.Println(*secondPosition)

	sort.Strings(names[:])

	fmt.Println(*secondPosition)

}