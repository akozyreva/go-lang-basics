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
}
