package main

import (
	"fmt"
	"strconv"
)

// func printPrice() {
// 	kayakPrice := 275.00
// 	kayakTax := kayakPrice * 0.2
// 	fmt.Println("Price : ", kayakPrice, "Tax : ", kayakTax)
// }

// func printPrice(product string, price float64, taxrate float64) {

// func printPrice(product string, price, taxrate float64) {
	
// 	taxAmount := price * taxrate

// func printPrice(product string, price, _ float64) {
	
// 	taxAmount := price * 0.2

// 	fmt.Println(product, "Price : ", price, "Tax : ", taxAmount)
// }

// func printSuppliers(product string, suppliers []string) {
// func printSuppliers(product string, suppliers ...string) {
// 		for index, supplier := range suppliers {
// 		fmt.Println("Product : ", product, "Suppliers : ", supplier, "Index : ", index)
// 	}

// }

// func swap(a int, b int) {
// 	fmt.Println("Before Swap : ", a, b)
// 	temp := a
// 	a = b
// 	b = temp
// 	fmt.Println("Swap Result : ", a, b)
// }

// func swap(a *int, b *int) {
// 	fmt.Println("Before Swap : ", *a, *b)
// 	temp := *a
// 	*a = *b
// 	*b = temp
// 	fmt.Println("Swap Result : ", *a, *b)
// }
	
// func shownum(a int, b int) (string) {
// 	return "Numbers are " + strconv.Itoa(a) + " and " + strconv.Itoa(b)
// 	// return "Hello World!"
// }

func showbignum(a int, b int) (k string) {
	defer fmt.Println("First defer call")

	if a > b {
		k = "Bigger number is " + strconv.Itoa(a)
	} else {
		k = "Bigger number is " + strconv.Itoa(b)
	}
	defer fmt.Println("Second defer call")


	return k
}

func main() {
	// fmt.Println("Hello functions")
	// printPrice()
	// printPrice("Kayak", 275, 0.2)

	// printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	// printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")

	// names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	// printSuppliers("Kayak", names...)

	// swap(1, 2)

	x, y := 10, 20

	// swap(&x, &y)

	// z := shownum(x, y)

	z := showbignum(x, y)

	fmt.Println(z)
}
