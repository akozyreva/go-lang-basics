package main

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var customerName string
	var label string = "Peter"

	quantity = 2
	customerName = "John Doe"
	/*
		customerName = 1
		error! as in Java, not possible to change type
	*/
	length, width = 1.2, 2.4
	fmt.Println(quantity)
	fmt.Println(customerName)
	fmt.Println(length, width)
	fmt.Println(label)

	var number int
	var name string
	var isFalse bool

	fmt.Println(number, name, isFalse) // 0, "", false

	// short form of assign of vars:
	// it's shorter!
	phoneNumber := "123-456-789"
	lastName := "Mark Levis"

	fmt.Println(lastName)
	fmt.Println(phoneNumber)
}
