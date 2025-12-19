// go run floats.go
// [71.8 56.2 89.5]s.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloats(fileName string) ([3]float64, error) {
	var floats [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // this obj is required for reading
	index := 0
	for scanner.Scan() {
		floats[index], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return floats, err // more easier, just return unlucky arr
		}
		index++
		// error can be here, but we don't read it here
	}
	err = file.Close()
	if err != nil {
		return [3]float64{}, err
	}
	if scanner.Err() != nil { // returns empty arr and Err()
		return [3]float64{}, scanner.Err()
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
