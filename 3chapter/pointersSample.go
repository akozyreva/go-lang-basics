package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)                 // prints address of amount in memory
	fmt.Println(reflect.TypeOf(amount))  // int
	fmt.Println(reflect.TypeOf(&amount)) // type of pointer => *int
	var myIntPointer *int                // init variable of pointer
	myIntPointer = &amount               // it refers to the same value as amount!!
	fmt.Println(myIntPointer)            // myIntPointer == &amount
	fmt.Println(*myIntPointer)           // 6. it shows the VALUE of int
	*myIntPointer = 10
	// it changes for both vars - amount and myIntPointer!
	fmt.Println(*myIntPointer)
	fmt.Println(amount)
}
