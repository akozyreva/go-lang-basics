package main

import (
	"fmt"
	"reflect"
)

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

	var length1 float64 = 1.2
	var width1 int = 2
	//fmt.Println("Area is", width1*length1) // error -> different types! float64!=int
	fmt.Println("Area is", width1*int(length1)) // all numbers are int, but it doesn't change length1 itself
	fmt.Println(reflect.TypeOf(length1))
}
