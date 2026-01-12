package main

import (
	"fmt"
	"math/rand"
)

type Stock struct {
	name  string
	price float32
}

func randomNumberGen(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

// Task implementation goes here
func (stock *Stock) updateStock() {
	change := randomNumberGen(-10000, 10000)
	stock.price += change
	stock.displayMarket()
}
func (stock *Stock) displayMarket() {
	fmt.Printf("Name: %s\n", stock.name)
	fmt.Printf("Price: %.2f", stock.price)
}
func main() {
	stockMarket := Stock{
		name:  "GOOG",
		price: 2313.50,
	}
	stockMarket.displayMarket()
	stockMarket.updateStock()

}
