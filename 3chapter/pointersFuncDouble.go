package main

import "fmt"

func double(amount *int) {
	*amount *= 2
}

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}
