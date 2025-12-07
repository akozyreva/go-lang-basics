package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello \nworld")
	fmt.Println("hello, \tworld")
	fmt.Println("hello, \"world")
	fmt.Println("hello, \\world")
	fmt.Println('Ж') // '' -> выводит Unicode значение
	fmt.Println(42)
	fmt.Println(3.1415)
	fmt.Println(5 < 2)

	// detect type of variable
	fmt.Println(reflect.TypeOf(45))
	fmt.Println(reflect.TypeOf(true))
}
