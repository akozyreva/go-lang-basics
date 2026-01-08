package main

import (
	"fmt"
)

func getUserInput[T any](inputValue T) {
	_, err := fmt.Scan(inputValue)
	for err != nil {
		fmt.Println("False value input! Try again, error is", err)
		_, err = fmt.Scan(inputValue)
	}
}

func main() {
	currencies := map[string]float32{
		"JPY": 130.2,
		"EUR": 0.95,
	}
	var dollarAmount float32
	var currency string
	fmt.Println("How many dollars do you have?")
	getUserInput(&dollarAmount)
	fmt.Print("Get currency: ")
	getUserInput(&currency)
	rate, isValid := currencies[currency]
	if isValid {
		fmt.Printf("Amount is: %.2f\n", rate*dollarAmount)
	} else {
		fmt.Println("Invalid currency")
	}
}
