package main

import "fmt"

type Cake struct {
	typeOfCake string
	weight     int // Weight in grams
}

func main() {

	// Checkpoint 1 code goes here
	cakes := []Cake{
		{typeOfCake: "Chocolate", weight: 1000},
		{typeOfCake: "Carrot", weight: 500},
		{typeOfCake: "Ice Cream", weight: 2000},
	}
	// Checkpoint 2 code goes here
	fmt.Println(cakes[0].weight)

	// Checkpoint 3 code goes here
	cakes[1].weight = 1500

	fmt.Println(cakes)

}
