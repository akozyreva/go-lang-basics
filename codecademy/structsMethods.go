package main

import "fmt"

type Triangle struct {
	height float32
	base   float32
}

// it's method of Triangle. it refers to Triangle. it's not function!!
// it's used to be OOP-like
func (triangle Triangle) area() float32 {
	return 0.5 * (triangle.height * triangle.base)
}

func main() {

	triangle := Triangle{10, 4}
	fmt.Println(triangle)

	// Call the function here
	area := triangle.area()
	fmt.Println(area)

}
