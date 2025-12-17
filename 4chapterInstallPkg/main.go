package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	number, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(number)
}
