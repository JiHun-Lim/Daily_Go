package main

import (
	"fmt"
	"time"
)


func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 4)
	for category, group := range data {
		go group.TotalPrice(category, channel)
		// storeTotal += group.TotalPrice(category)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("-- Starting to receive from channel")
	for i := 0; i< len(data); i++ {
		fmt.Println("-- channel read pending", len(channel), "items in buffer, siz", cap(channel))
		value := <- channel
		fmt.Println("-- channel read complete", value)
		storeTotal += value
		time.Sleep(time.Second)
		// storeTotal += <- channel
	}

	fmt.Println("Total  : ", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		// fmt.Println(category, "product : ", p.Name)
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "channel sending", ToCurrency(total))
	resultChannel <- total
	fmt.Println(category, "channel send complete")
}