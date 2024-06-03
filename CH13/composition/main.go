package main

import (
	"fmt"
	"composition/store"
)

func main() {

	// kayak := store.NewProduct("Kayak", "Watersports", 275)
	// lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	// for _, p := range []*store.Product {kayak, lifejacket} {
	// 	fmt.Println("Name:", p.Name, "Category:", p.category, "Price:", p.Price(0.2))
	// }

	// boats := []*store.Boat {
	// 	store.NewBoat("Kayak", 275, 1, false),
	// 	store.NewBoat("Canoe", 400, 3, false),
	// 	store.NewBoat("Tender", 650.25, 2, true),
	// }

	// for _, b := range boats {
	// 	fmt.Println(b.Product.Name, b.Name)
	// 	fmt.Println(b.Name, b.Price(0.2))
	// }

	// rentals := []*store.RentalBoat {
	// 	store.NewRentalBoat("Kayak", 275, 1, false, false),
	// 	store.NewRentalBoat("Canoe", 400, 3, false, true),
	// 	store.NewRentalBoat("Tender", 650.25, 2, true, true),
	// }

	// for _, r := range rentals {
	// 	fmt.Println(r.Name, r.Price(0.2))
	// }

	rentals := []*store.RentalBoat {
		store.NewRentalBoat("Kayak", 275, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Canoe", 400, 3, false, true, "Bob", "Alice"),
		store.NewRentalBoat("Tender", 650.25, 2, true, true, "Dora", "Charlie"),
	}

	for _, r := range rentals {
		fmt.Println(r.Name, r.Price(0.2), r.Captain)
	}



}