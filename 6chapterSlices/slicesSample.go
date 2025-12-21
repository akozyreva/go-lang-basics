package main

import "fmt"

func main() {
	// slices sample
	var mySlice []string        // init of slice
	mySlice = make([]string, 7) // make is required to create slice!
	fmt.Println(mySlice)

	// short form
	mySlice1 := make([]string, 7)
	fmt.Println(mySlice1)

	// short form with init
	notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println("notes:", notes)

}
