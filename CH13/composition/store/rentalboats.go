package store

type RentalBoat struct {
	*Boat
	*Crew
	IncludeCrew bool
}

type Crew struct {
	Captain, FirstOfficer string
}

// func NewRentalBoat(name string, price float, capacity int, motorized, crewed bool) *RentalBoat {
// 	return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed}
// }

func NewRentalBoat(name string, price float, capacity int, motorized, crewed bool, captain, firstofficer string) *RentalBoat {
	return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed, &Crew{capacity, firstofficer}}
}