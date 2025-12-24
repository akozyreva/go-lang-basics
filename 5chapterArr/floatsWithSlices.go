// go run floatsWithSlices.go
// [71.8 56.2 89.5]s.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloats(fileName string) ([]float64, error) {
	floats := make([]float64, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // this obj is required for reading
	index := 0
	for scanner.Scan() {
		data, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err // just return nil, because it's obvious, we don't have normal result
		}
		floats = append(floats, data)
		index++
		// error can be here, but we don't read it here
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil { // returns empty arr and Err()
		return nil, scanner.Err()
	}
	return floats, nil

}

func main() {
	numbers, err := getFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)

}
