package main

import (
	"fmt"
)

func severalInt(numbers ...int) { // ...int - it's actually slice!
	fmt.Println(numbers)
}

func main() {
	severalInt(1)
	severalInt(1, 2, 3)
	slice := []int{3, 2, 1}
	severalInt(slice...) // this syntax is used to unpack slice
	//severalInt(slice) -> err, because we're waiting for values, not for slice
}
