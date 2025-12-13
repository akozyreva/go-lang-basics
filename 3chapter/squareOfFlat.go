package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

// it's not an error not to repeat line
func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func calcArea(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f litres neeeded\n", area/10.0)
}

func main() {
	var width, height float64
	width = 4.2
	height = 3.0
	calcArea(width, height)
	sayHi()
}
