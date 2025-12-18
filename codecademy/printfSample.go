package main

import "fmt"

func main() {
	animal1 := 666
	animal2 := "dog"

	// %v refers to anything! to number or string
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)

	floatExample := 1.75
	// it prints type of var!!
	fmt.Printf("Working with a %T.", floatExample)
	//Working with a float64.

	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	meditation := fmt.Sprintln(step1, step2) // it's format string and return string value
	fmt.Println(meditation)

}
