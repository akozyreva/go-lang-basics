package main

import "fmt"

func inc(x *int) {
	*x++
	// I'm not able to write x+=5, because *x refers that it's pointer
}

func swap(x, y *int) {
	// поменяйте значения местами
	*x, *y = *y, *x
}

func main() {
	a := 5
	inc(&a)
	fmt.Println(a) // должно напечатать 6

	a, b := 3, 7
	swap(&a, &b)
	fmt.Println(a, b) // 7 3
}
