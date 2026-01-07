package main

import "fmt"

func main() {
	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println("myMap:", myMap)

	words := map[string]string{"de": "Hallo"}
	fmt.Println(words["de"])
	fmt.Println(words["en"]) // key doesn't exist, but it doesn't show an error, it shows empty value
	// it's zero value philosophy - minimum of errors
	fmt.Println("words:", words)
	numberOfRooms := make(map[string]int)
	numberOfRooms["Mike"]++ // it returns 1! it's very convenient for counting
	// removing key as in python
	delete(numberOfRooms, "Mike") // empty map
	fmt.Println("numberOfRooms:", numberOfRooms)
}
