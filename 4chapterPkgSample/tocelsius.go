package main

import (
	"4chapterPkgSample/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit:")
	farenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (farenheit - 32) * 5 / 9
	fmt.Println(celsius)
}
