package main

import (
	"errors"
	"fmt"
	"log"
)

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

func getCalcArea(width float64, height float64) (float64, error) {
	var err error
	if height < 0 {
		err = errors.New("height can't be negative")
		// area not counted, err returns
		// 0 - it's default value of float64. I have to return at least it
		return 0, err
	}
	if width < 0 {
		err = errors.New("width can't be negative")
		return 0, err
	}

	area := width * height
	fmt.Printf("%.2f litres neeeded\n", area/10.0)
	// nil means that there are no errors
	return area, nil
}

func main() {
	var width, height float64
	width = 4.2
	height = 3.0
	calcArea(width, height)
	sayHi()
	area, err := getCalcArea(width, height)
	// simulate error
	area, err = getCalcArea(width, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(area)
}
