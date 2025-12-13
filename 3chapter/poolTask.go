package main

import "fmt"

func main() {
	var myInt int
	var myPointer *int
	myInt = 42
	myPointer = &myInt      // set pointer to var myInt
	fmt.Println(*myPointer) // show the value on which refers pointer
}
