// task is the same as in floatsWithSlices.go
// but you need to take numbers from keyboard
// [71.8 56.2 89.5]s.go
// go run floatsWithSlicesWithArgs.go 71.8 56.2 89.5
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloats(arrArgs []string) ([]float64, error) {
	var floats []float64
	for _, el := range arrArgs {
		data, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return nil, err // just return nil, because it's obvious, we don't have normal result
		}
		floats = append(floats, data)
	}
	return floats, nil

}

func main() {
	fmt.Println(os.Args[1:]) // from 1 element, because 0el is path to file
	arr := os.Args[1:]
	numbers, err := getFloats(arr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)

}
