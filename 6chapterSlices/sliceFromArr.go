package main

import "fmt"

func main() {
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr)
	slice := arr[2:5] // it's also slice! it doesn't create copy!!!
	fmt.Println(slice)
	slice[0] = 100
	fmt.Println(slice) // 1 el is 100!
	fmt.Println(arr)   // 3 element is 100 as well
	// append value to slice - at that moment new arr is created with more memory
	// slice has access to it
	slice = append(slice, 1000)
	fmt.Println(slice) // new array!
	fmt.Println(arr)   // old array!

	// init empty slice
	var emptySlice []string // if no len, it's slice, no array!
	fmt.Println(emptySlice)
}
